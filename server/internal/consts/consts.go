package consts

const MySwaggerUITemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<meta name="description" content="api文档"/>
	<title>api文档</title>
	<link rel="stylesheet" href="/api/docs.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="/api/docs.js"></script>
<script>
	window.onload = () => {
		window.ui = SwaggerUIBundle({
			url:    '{SwaggerUIDocUrl}',
			dom_id: '#swagger-ui',
		});
	};
</script>
</body>
</html>
`
