package htmltpl

const TagTpl = `<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>{{ .Title }}</title>
    <style type="text/css">
		body{font-family:-apple-system,"Helvetica Neue",Helvetica,Arial,"PingFang SC","Hiragino Sans GB","WenQuanYi Micro Hei","Microsoft Yahei",sans-serif;font-size:14px;line-height:1.42858;color:#333;background-color:#fff;-webkit-font-smoothing:antialiased}
		pre{background-color:#e9ecef;padding:1em;border-radius:.3em;overflow:auto;}
		a{text-decoration:none}
		a:hover{text-decoration:underline}
		.clearboth{clear:both}
		.wrapper{margin-right:auto;margin-left:auto;padding-left:15px;padding-right:15px}
		.header{margin:1em 0 0 0;padding-bottom:1em;border-bottom:1px solid #eee}
		.header .logo{float:left;font-size:1.2em}
		.header .cates{float:right;margin-top:0.1em}
		.header .cates ul{margin:0;padding:0}
		.header .cates ul li{list-style:none;float:left;margin-right:.8em}
		.header .cates ul li:last-child{margin-right:0}
		.container{min-height:86%;border-bottom:1px solid #eef}
		.container ul.catelogs{margin:0;padding:0}
		.container .sidebar{box-sizing:border-box;width:16.66667%;float:left;padding-right:1em}
		.container .sidebar ul{margin:0;padding:0}
		.container ul.catelogs li{list-style:none;display:inline-block;margin-right:.6em;font-size:1.2em}
		.container .sidebar li{list-style:none}
		.container .sidebar ul.inlineblock li{display:inline-block;margin-right:.6em}
		.container .sidebar .item h3{border-top:1px solid #eee}
		.container .content{float:left;box-sizing:border-box;padding-left:1em;width:83.33333%}
		.footer{margin-top:0.5em;text-align:center}

		@media(min-width:768px){.wrapper{width:750px}}
		@media(min-width:992px){.wrapper{width:970px}}
	</style>
</head>
<body>
<div class="wrapper">
{{ template "header" .Header }}
{{ template "container" .Container }}
{{ template "footer" .Footer }}
</div>
</body>
</html>

{{ define "header" }}
<div class="header">
    <span class="logo"><a href="{{ rooturl }}/index.html">{{ .LogoName }}</a></span>
    <div class="cates">
        <ul>
            {{ range .Cates -}}
            <li><a href="{{ rooturl }}/{{ . }}/index.html">{{ . }}</a></li>
            {{ end }}
        </ul>
    </div>
    <div class="clearboth"></div>
</div>
{{ end }}

{{ define "sidebar" }}
<div class="sidebar">
    <h2>{{ .CurCate }}</h2>
    <div class="item">
        <h3>Articles: </h3>
        <ul>
        {{ range .Articles -}}
            <li><a href="{{ rooturl }}/{{ $.CurCate }}/{{ . }}.html">{{ . }}</a></li>
        {{ end }}
        </ul>
    </div>
    <div class="item">
        <h3>Tags: </h3>
        <ul>
            {{ range .Tags -}}
            <li><a href="{{ rooturl }}/{{ $.CurCate }}/tag/{{ . }}.html">{{ . }}</a></li>
            {{ end }}
        </ul>
    </div>
    <div class="clearboth"></div>
</div>
{{ end }}

{{ define "content" }}
<div class="content">
    <h1>Tag: {{ .TagName }}</h1>
    <ul>
        {{ range .Articles -}}
        <li><a href="{{ rooturl }}/{{ $.CurCate }}/{{ . }}.html">{{ . }}</a></li>
        {{ end }}
    </ul>
</div>
{{ end }}

{{ define "container" }}
<div class="container">
    {{ template "sidebar" .Sidebar }}
    {{ template "content" .Content }}
    <div class="clearboth"></div>
</div>
{{ end }}

{{ define "footer" }}
<div class="footer">
{{ . }}
</div>
{{ end }}
`
