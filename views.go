package main

const EDIT_PAGE_TEMPLATE = `
{{ define "content" }}
<div class="page-header">
  <h1>Edit page</h1>
</div>
<form action="/pages/{{.Name}}" method="post">
  <div class="form-group"><input type="text" name="title" value="{{.Title}}" class="form-control"></div>
  <div class="form-group"><textarea name="content" rows="25" class="form-control">{{.Markdown}}</textarea></div>
  <div class="form-group"><input type="text" name="message" class="form-control"></div>
  <div class="form-group">
    <input type="submit" value="Save" class="btn btn-primary">
    <a href="/" class="btn btn-default">Cancel</a>
  </div>
</form>
{{ end }}
`

const NEW_PAGE_TEMPLATE = `
{{ define "content" }}
  <div class="page-header">
    <h1>Create a new page</h1>
  </div>
  <form action="/pages" method="post">
    <input type="hidden" name="name">
    <div class="form-group">
      <input type="text" name="title" class="form-control">
    </div>
    <div class="form-group">
      <textarea name="content" rows="25" class="form-control"></textarea>
    </div>
    <div class="form-group">
      <input type="submit" value="Save" class="btn btn-primary">
      <a href="/" class="btn btn-default">Cancel</a>
    </div>
  </form>
{{ end }}
`

const NEW_NAMED_PAGE_TEMPLATE = `
{{ define "content" }}
  <div class="page-header">
    <h1>Create a new page</h1>
  </div>
  <form action="/pages" method="post">
    <div class="form-group">
      <input type="text" name="title" value="{{.}}" class="form-control">
    </div>
    <div class="form-group">
      <textarea name="content" rows="25" class="form-control"></textarea>
    </div>
    <div class="form-group">
      <input type="submit" value="Save" class="btn btn-primary">
      <a href="/" class="btn btn-default">Cancel</a>
    </div>
  </form>
{{ end }}
`

const LOGIN_TEMPLATE = `
{{ define "content" }}
<div id="content">
  <a href="">Login with github</a>
</div>
{{ end }}
`

const SHOW_WIKI_PAGE_TEMPLATE = `
{{ define "content" }}
<div id="content">
  <div class="page-header"><h1>{{.Title}}</h1></div>
  <div class="btn-group" role="group" aria-label="...">
    <a href="/pages/{{.Name}}/edit" class="btn btn-default">Edit</a>
    <a href="/wiki/{{.Name}}/history" class="btn btn-default">History</a>
    <a href="/wiki" class="btn btn-default">Pages</a>
    <a href="/pages/new" class="btn btn-default">New</a>
  </div>
  <div>{{.Markup}}</div>
</div>
{{ end }}
`

const ERROR_PAGE_TEMPLATE = `
{{ define "content" }}
<div id="content">
  <div class="alert alert-danger" role="alert"> {{.}} </div>
</div>
{{ end }}
`
