{{template "header.html" .}}

<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css"
  integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

<div class="container-fluid">
  <form class="form-horizontal" id="">
    <div class="form-group">
      <label for="inputEmail3" class="col-sm-2 control-label">标题</label>
      <div class="col-sm-10">
        <input type="text" name='title' class="form-control" id="inputEmail3" placeholder="标题">
      </div>
    </div>
    <div class="form-group">
      <label for="inputPassword3" class="col-sm-2 control-label">作者</label>
      <div class="col-sm-10">
        <input type="text" name='author' class="form-control" id="inputPassword3" placeholder="作者">
      </div>
    </div>
    <div class="form-group">
      <label for="inputPassword3" class="col-sm-2 control-label">简介</label>
      <div class="col-sm-10">
        <textarea name='desc' class="form-control" rows="3"></textarea>
      </div>
    </div>
    <div class="form-group">
      <label for="inputPassword3" class="col-sm-2 control-label">是否显示</label>
      <div class="col-sm-10">
        <label class="radio-inline">
          <input type="radio" name="status" id="inlineRadio1" value="1"> 显示
        </label>
        <label class="radio-inline">
          <input type="radio" name="status" id="inlineRadio2" value="2"> 隐藏
        </label>
      </div>
    </div>
    <div class="form-group">
      <label for="inputPassword3" class="col-sm-2 control-label">分类</label>
      <div class="col-sm-10">
        <select class="form-control" name='class'>
          <option>玄幻</option>
          <option>言情</option>
          <option>科技</option>
          <option>鬼怪</option>
          <option>色情</option>
        </select>
      </div>
    </div>

    <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <button type="button" onclick="add()" class="btn btn-primary">提交</button>
      </div>
    </div>
  </form>
</div>
<script>
  function add() {
    $('.form-horizontal').ajaxSubmit({
      url: '//127.0.0.1:8080/novel/add',
      type: 'post',
      //dataType:'json',
      success: function (data) {
        if (data.Code == 1) {
          window.location.href = "//127.0.0.1:8080/novel"
        }
      }
    })
  }
</script>