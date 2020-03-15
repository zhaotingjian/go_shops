<html>

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>crm管理系统</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css"
  integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
  <script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.js"></script>
</head>
<style>
  .form {
    background: #ccc;
    width: 500px;
    margin-top: 240px;
    margin-left: 480px;
  }
</style>
<body>
  <div class="container-fluid">
    <form class="form-horizontal" style="margin-top: 16%;margin-left:35%;width: 410px;">
      <div class="form-group">
        <label for="inputEmail3" class="col-sm-2 control-label">用户名</label>
        <div class="col-sm-10">
          <input type="text" name="username" class="form-control" id="username" placeholder="Username">
        </div>
      </div>
      <div class="form-group">
        <label for="inputPassword3" class="col-sm-2 control-label">密码</label>
        <div class="col-sm-10">
          <input type="password" name="password" class="form-control" id="password" placeholder="Password">
        </div>
      </div>
      <div class="form-group">
        <div class="col-sm-offset-2 col-sm-10">
          <div class="checkbox">
            <label>
              <input type="checkbox"> Remember me
            </label>
          </div>
        </div>
      </div>
      <div class="form-group">
        <div class="col-sm-offset-2 col-sm-10">
          <button type="button" id="login" class="btn btn-default">登入</button>
        </div>
      </div>
    </form>
  </div>
</body>
</html>
<script>
  $(function () {
    $('#login').bind('click', login);
  })
  function login() {
    var username = $("#username").val();
    var password = $("#password").val();
    var url = "//127.0.0.1:8080/login";
    $.ajax(
      {
        url: url,
        type: "post",
        data: { username: username, password: password },
        dataType: "json",
        success: function (data) 
        {
          alert(data.Msg)
          if (data.Code == 1) 
          {
            window.location="//127.0.0.1:8080/main";
          }
        }
      }
    );
  }
</script>