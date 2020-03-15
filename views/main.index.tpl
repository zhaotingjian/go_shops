<html>

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
  <title>Bootstrap 101 Template</title>
  <script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.js"></script>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css"
    integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
  <script src=""></script>
</head>

<body>
  <div class="container-fluid">
    <div class="row">
      <div class="col-xs-12" style="height:70px;background-color: bisque;">
        .col-xs-9
      </div>
      <div class="col-xs-2" style="height:690px;background-color:blue;">
        <div class="list-group" style="width:100%;text-align: center;">
          <a href="#" class="list-group-item first_menu" target="frame">首页</a>
          <div>
            <a href="/novel" class="list-group-item first_menu" target="frame">小说管理</a>
            <div class="list-group second_menu" style="width:100%;text-align: center;margin-bottom: 0px;display: none;">
              <a href="/novel" class="list-group-item active" target="frame">小说列表</a>
              <a href="/novel/log" class="list-group-item active" target="frame">操作日志</a>
            </div>
          </div>
          <div>
            <a href="/employee" class="list-group-item first_menu" target="frame">员工管理</a>
            <div class="list-group second_menu" style="width:100%;text-align: center;margin-bottom: 0px;display: none;">
              <a href="/employee" class="list-group-item active" target="frame">员工列表</a>
              <a href="/employee/login_log" class="list-group-item active" target="frame">登入日志</a>
              <a href="/employee/oprate_log" class="list-group-item active" target="frame">操作日志</a>
            </div>
          </div>
          
        </div>
      </div>
      <div class="col-xs-10" style="height:690px;background-color:brown;">
        <iframe name="frame" src="" style="width:100%;height:100%"></iframe>
      </div>
    </div>
  </div>
</body>
<script>
  $(function(){
    $("a.first_menu").bind('click',function(){
      $('div.second_menu').hide()
      _this=$(this);
      _second_menu=_this.parent('div').find('.second_menu');
      if(_second_menu.css('display')=='none'){
         _second_menu.show()
      }
      
    })
  })
</script>
</html>