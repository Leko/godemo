{{template "header.tpl" .}}

<h1>Go demo application</h1>

<a href="/login">Login</a> or <a href="/register">Register</a>

{{if .user.ID}}
<div class="alert alert-warning">
	<dl>
		<dt>Email</dt>
		<dd>{{.user.Email}}</dd>

		<dt>Password(Hashed)</dt>
		<dd>{{.user.Password}}</dd>

		<dt>CreatedAt</dt>
		<dd>{{.user.CreatedAt}}</dd>
	</dl>

	<a href="/logout" class="btn btn-primary">Log out</a>
</div>
{{end}}

{{template "footer.tpl" .}}
