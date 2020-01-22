<html>

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Bootstrap 101 Template</title>
  <!-- <link href="../../static/css/index.css" rel="stylesheet"> -->
  <script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.js"></script>
  <script src=""></script>
</head>
<style>
  .form {
    background: #ccc;
    width: 500px;
    margin-top: 240px;
    margin-left: 480px;
  }
</style>

<body style="background: #ffffff;">
  <div class="form" id="form">
    <div>用户名<input type="text" name="username" id="username" /></div>
    <div>密码<input type="text" name="password" id="password" /></div>
    <div>
      <button class="login" id="login">登录</button>
      <button class="cancel" id="cancel">取消</button>
    </div>
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
    var url = "http://127.0.0.1:8080/login";
    $.ajax({
      url: url,
      type: "post",
      data: { username: username, password: password },
      dataType: "json",
      success: function (data) {
        alert(data.Msg)
        if (data.Code == 1) {
          window.location="http://127.0.0.1:8080/index"
        }
      }
    });
  }
</script>