<!DOCTYPE html>

<html>
<head>
  <title>Box Portal</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet"  type="text/css" media="screen" href="/static/css/main.css"/>

  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

  <!-- Optional theme -->
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

  <script src="//cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script>

  <!-- Latest compiled and minified JavaScript -->
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
</head>
  <header class="navbar navbar-default navbar-static-top">
    <div class="container">
      <div class="navbar-header">
        <a class="navbar-brand" href="#">Box Portal</a>       
      </div>
      <p class="navbar-text navbar-right">
        <a href="/login">Sign In</a>
      </p>
      <p class="navbar-text navbar-right">Register</p>     
    </div>
  </header>

<div class="container-fluid" id="header-image">
</div>

<div id="main" class="container">
  <div class="row">
    <div class="col-md-7">
      <div class="container">
        <h2>Register</h2>
        <p>Create an account to access box.xzdbd.com.</p>
        <h2>What is Box?</h2>
        <p>Box is a file storage and sharing portal.</p>
        <button type="button" class="btn btn-primary">Register a Box ID</button>
      </div>
    </div>
    <div class="col-md-5">
      <div id="signin">
        <h2>Sign In</h2>
        <p>Sign in to Box</p>
        <form enctype="application/x-www-form-urlencoded" method="post">
          <div class="form-group">
            <label for="username">USERNAME</label>
            <input type="text" class="form-control" name="username" id="usernameInput" placeholder="username">
          </div>
          <div class="form-group">
            <label for="username">PASSWORD</label>
            <input type="password" class="form-control" name="password" id="passwordInput" placeholder="password">
          </div>
          <button type="submit" class="btn btn-lg btn-primary">Sign In</button>
        </form>
      </div>
    </div>
  </div>
</div>

<div id="footer" class="container">
  <ul id="footer-list" class="list-inline">
    <li>ABOUT BOX</li>
    <li>DOCS</li>
    <li>POLICY</li>
    <li>CONTACT</li>
  </ul>  
  <div id="footer-siteinfo">
    <p>
      Copyright © 2015-2016 xzdbd.com. All rights reserved.
    </p>
  </div>
</div>
</html>
