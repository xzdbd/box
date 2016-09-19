<!DOCTYPE html>

<html>
<head>
  <title>Login - Box Portal</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
  <link rel="stylesheet"  type="text/css" media="screen" href="/static/css/main.css"/>
  <link href="/static/css/semantic/semantic.min.css" rel="stylesheet">
  
  <script src="//cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script>
  <script src="/static/js/semantic/semantic.min.js"></script>
  <style type="text/css">
    body {
      background-color: #DADADA;
    }
    body > .grid {
      height: 100%;
    }
    .image {
      margin-top: -100px;
    }
    .column {
      max-width: 450px;
    }
  </style>
  <script>
  $(document)
    .ready(function() {
      $('.ui.form')
        .form({
          fields: {
            username: {
              identifier  : 'username',
              rules: [
                {
                  type   : 'empty',
                  prompt : 'Please enter your username'
                } 
              ]
            },
            password: {
              identifier  : 'password',
              rules: [
                {
                  type   : 'empty',
                  prompt : 'Please enter your password'
                }
              ]
            }
          }
        })
      ;
    })
  ;
  </script>
  <style id="style-1-cropbar-clipper">/* Copyright 2014 Evernote Corporation. All rights reserved. */
  .en-markup-crop-options {
      top: 18px !important;
      left: 50% !important;
      margin-left: -100px !important;
      width: 200px !important;
      border: 2px rgba(255,255,255,.38) solid !important;
      border-radius: 4px !important;
  }

  .en-markup-crop-options div div:first-of-type {
      margin-left: 0px !important;
}
</style>
</head>
<body>
  <div class="ui middle aligned center aligned grid">
    <div class="column">
      <h2 class="ui teal image header">
        <img src="/static/img/logo.jpeg" class="image">
        <div class="content">
          Sign In to <em>Box</em>
        </div>
      </h2>
      <form class="ui large form" method="post" enctype="application/x-www-form-urlencoded">
        <div class="ui stacked segment">
          <div class="field">
            <div class="ui left icon input">
              <i class="user icon"></i>
              <input type="text" name="username" placeholder="Username">
            </div>
          </div>
          <div class="field">
            <div class="ui left icon input">
              <i class="lock icon"></i>
              <input type="password" name="password" placeholder="Password">
            </div>
          </div>
          <div class="ui fluid large teal submit button">Sign In</div>
        </div>
        <div class="ui error message" style="{{if .isLoginFail}}display: block{{else}}{{end}}">
          {{if .isLoginFail}}
          <ul class="list">
            <li>
              Your username and password did not match. Please try again.
            </li>
          </ul>
          {{else}}
          {{end}}
        </div>
      </form>
      <div class="ui message">
      New to us? <a href="#">Sign Up</a>
      </div>
    </div>
  </div>
</body>
<div></div>
</html>
