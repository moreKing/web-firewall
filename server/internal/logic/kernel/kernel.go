package kernel

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"os/exec"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/service"
	"strings"
)

type sKernel struct{}

const id = 11

func init() {
	service.RegisterKernel(New())
}

func New() service.IKernel {
	s := &sKernel{}
	// 初始化系统参数
	get, err := s.Get(context.Background())
	if err != nil {
		panic(err)
	}
	err = s.Set(context.Background(), get)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *sKernel) Set(ctx context.Context, m *model.Kernel) error {
	_, err := dao.SystemConf.Ctx(ctx).Where(dao.SystemConf.Columns().Id, id).Update(do.SystemConf{
		Config: m,
	})
	if err != nil {
		return err
	}

	// 对数据进行初始化配置
	if m.Forward {
		if err := update("net.ipv4.ip_forward", "1"); err != nil {
			return err
		}
		if err := update("net.ipv6.conf.all.forwarding", "1"); err != nil {
			return err
		}

	} else {
		if err := update("net.ipv4.ip_forward", "0"); err != nil {
			return err
		}
		if err := update("net.ipv6.conf.all.forwarding", "0"); err != nil {
			return err
		}
	}

	// 不对结果进行判断 直接判断运行文件结果是否是预期结果
	//if err := sysctl(); err != nil {
	//	return err
	//}

	return nil
}

func (s *sKernel) Get(ctx context.Context) (m *model.Kernel, err error) {
	var conf model.SystemConfKernel
	err = dao.SystemConf.Ctx(ctx).Where(dao.SystemConf.Columns().Id, id).Scan(&conf)
	if err != nil {
		return nil, err
	}
	g.Log().Debug(ctx, conf)
	return &conf.Config, nil
}

func update(key, value string) error {
	if strings.TrimSpace(key) == "" || strings.TrimSpace(value) == "" {
		return nil
	}

	file, err := os.ReadFile("/etc/sysctl.conf")
	// 读取 /etc/sysctl.conf 文件

	list := strings.Split(string(file), "\n")
	g.Log().Debug(context.Background(), list)
	for i, line := range list {
		if strings.HasPrefix(strings.TrimSpace(line), "#") || strings.TrimSpace(line) == "" {
			continue
		}

		item := strings.Split(line, "=")
		if len(item) != 2 {
			continue
		}
		if strings.TrimSpace(item[0]) == key {
			list[i] = fmt.Sprintf("%s=%s", key, value)
			goto WRITE
		}

	}

	list = append(list, fmt.Sprintf("%s=%s", key, value), "\n")

	// 写入文件
WRITE:
	g.Log().Debug(context.Background(), strings.Join(list, "\n"))

	file2, err := os.OpenFile("/etc/sysctl.conf", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer func(file2 *os.File) {
		_ = file2.Close()
	}(file2)

	_, err = file2.WriteString(strings.Join(list, "\n"))
	if err != nil {
		return err
	}
	// 不进行加载文件 直接写入运行文件中
	err = exec.Command("sh", "-c", fmt.Sprintf("echo %s > /proc/sys/%s", value, strings.Replace(key, ".", "/", -1))).Run()
	if err != nil {
		return err
	}

	// 判断是否与预期一致

	output, err := exec.Command("cat", fmt.Sprintf("/proc/sys/%s", strings.Replace(key, ".", "/", -1))).CombinedOutput()
	if err != nil {
		return err
	}

	if strings.TrimSpace(string(output)) != strings.TrimSpace(value) {
		return errors.New(fmt.Sprintf("参数 %s 修改失败,预期%s, 实际%s ", key, value, string(output)))
	}

	return nil
}

// 配置生效
func sysctl() error {
	err := exec.Command("sysctl", "-p").Run()
	if err != nil {
		return err
	}

	return nil
}
