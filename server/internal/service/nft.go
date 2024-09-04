// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/model"
)

type (
	INft interface {
		Add(ctx context.Context, rulesets *model.Rulesets, add bool) error
		// func (s *sNft) Replace() {
		//
		// }
		//
		// func (s *sNft) ChangePosition() {
		//
		// }
		//
		// func (s *sNft) Delete() {
		//
		// }
		GetChainList(chain int) *[]model.Rulesets
		Delete(ctx context.Context, id int64) error
		Replace(ctx context.Context, rule *model.Rulesets) error
		UpdatePosition(ctx context.Context, rid, did int64, add bool) error
	}
)

var (
	localNft INft
)

func Nft() INft {
	if localNft == nil {
		panic("implement not found for interface INft, forgot register?")
	}
	return localNft
}

func RegisterNft(i INft) {
	localNft = i
}
