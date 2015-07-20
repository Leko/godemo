{{template "header.tpl" .}}

<div class="row">
  <div class="col-md-4 col-md-offset-4">
    {{if .new }}
    <h1>Register to demoapp</h1>
    <form action="/users/create" method="POST">
    {{else}}
    <h1>Login to demoapp</h1>
    <form action="/authenticate" method="POST">
    {{end}}
      <input type="hidden" name="csrf_token" value="{{ .csrfToken }}">
      <div class="form-group">
        <label for="form-email" class="form-label">Email</label>
        <input type="email" id="form-email" class="form-control" name="email">
      </div>
      <div class="form-group">
        <label for="form-password" class="form-label">Password</label>
        <input type="password" id="form-password" class="form-control" name="password">
      </div>
      <div class="form-group">
        {{if .new }}
        <input type="submit" class="btn btn-primary" value="Register">
        or <a href="/login">Login</a>
        {{else}}
        <input type="submit" class="btn btn-primary" value="Login">
        or <a href="/register">Register</a>
        {{end}}
      </div>
    </form>
  </div>
</div>

{{template "footer.tpl" .}}
