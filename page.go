package main

var pageTemplate = `<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>This site has moved</title>
    <style>
        body {
            background-color: #eaeaea;
            letter-spacing: 0.05rem;
            font-family: sans-serif;
            padding: 2rem;
        }
        h3 {
            padding-bottom: 2rem;
        }
    </style>
</head>
<body>
<h1>This site has moved</h1>
{{ if .AdditionalMessage }}
    <h3>{{.AdditionalMessage}}</h3>
{{ else }}
    <h3><em>{{.OldHost}}</em> has changed to <em>{{.NewHost}}</em></h3>
{{ end }}
<h3>Please update any bookmarks and links you have to the new url:
<a href="{{.NewURL}}">{{.NewURL}}</a>
</h3>
{{ if .RedirectEndDate }}
    <p><em>This message will be available until {{.RedirectEndDate}}</em></p>
{{ end }}
{{ if .MoreInfoURL }}
<p>Learn more about this change at: <a target="_blank" href="{{.MoreInfoURL}}">{{.MoreInfoURL}}</a></p>
{{ end }}
</body>
</html>`
