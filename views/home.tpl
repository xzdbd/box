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
        <tr>
          <td><a href="/disk/home">filename111.txt</a></td>
          <td>2016-08-09 15:00</td>
          <td><button class="ui primary button" onclick="window.location.href='./share?objectName=objectname&fileName=filename'">共享</button></td>         
        </tr>
        {{str2html .UserObjects}}
      </tbody>
    </table>  
  </div>
</body>
<div></div>
</html>
