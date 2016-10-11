<!DOCTYPE html>

<html>
<head>
  <title>Home - Box Portal</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
  <link rel="stylesheet"  type="text/css" media="screen" href="/static/css/main.css"/>
  <link href="/static/css/semantic/semantic.min.css" rel="stylesheet">
  
  <script src="//cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script>
  <script src="/static/js/semantic/semantic.min.js"></script>
  <script type="text/javascript">
    $(function (){
      $('.message .close')
        .on('click', function() {
          $(this)
            .closest('.message')
            .transition('fade')
          ;
        })
      ;
    })
  </script>
</head>
<body>
  <div class="ui container">
    <div class="ui menu">
      <div class="header item">
        Box
      </div>
      <a class="item" href="/disk/home">Home</a>
      <div class="right menu">
        <a class="item" href="/disk/home?action=logout">注销</a>
      </div>
    </div>
    <h1>文件列表</h1>
    <div>
      {{str2html .ShareMessage}}
    </div>
    <table class="ui selectable olive table">
      <thead>
        <tr>
          <th class="nine wide">名称</th>
          <th class="four wide">最后修改时间</th>
          <th class="three wide">共享</th>
        </tr>
      </thead>
      <tbody>
        {{str2html .UserObjects}}
      </tbody>
    </table>  
  </div>
  <div class="ui vertical center aligned segment">
    <div class="ui blue label">2016 © xzdbd All Rights Reserved.</div>
    <div class="ui orange label">Version: 0.0.1</div>
  </div>
</body>
<div></div>
</html>
