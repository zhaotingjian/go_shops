{{template "header.html" .}}
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css"
  integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
<script src="/static/js/bootstrap-paginator.min.js"></script>

<div class="container-fluid">
  <!-- 搜索框 -->
  <div class="panel panel-default">
    <form class="form-inline" action="/novel" method="get" id="search_data">
      <div class="form-group">
        <label for="exampleInputName2">ID</label>
        <input type="text" class="form-control" name="id" id="exampleInputName2" value="{{.Search.id}}" />
      </div>
      <div class="form-group">
        <label for="exampleInputName2">用户名</label>
        <input type="text" class="form-control" name="title" id="exampleInputName2" value="{{.Search.title}}" />
      </div>
      <div class="form-group">
        <label for="exampleInputName2">真实姓名</label>
        <input type="text" class="form-control" name="author" id="exampleInputName2" value="{{.Search.author}}" />
      </div>
      <div class="form-group">
        <label for="exampleInputName2">昵称</label>
        <input type="text" class="form-control" name="status" id="exampleInputName2" value="{{.Search.status}}" />
      </div>
      <div class="form-group">
        <label for="exampleInputName2">性别</label>
        <input type="text" class="form-control" name="creat_time" id="exampleInputName2" value="{{.Search.creat_time}}" />
      </div>
      <div class="form-group">
        <label for="exampleInputEmail2">年龄</label>
        <input type="email" class="form-control" name="edit_time" id="exampleInputEmail2" value="{{.Search.edit_time}}" />
      </div>
      <div class="form-group">
        <label for="exampleInputEmail2">手机号</label>
        <input type="email" class="form-control" name="edit_time" id="exampleInputEmail2" value="{{.Search.edit_time}}" />
      </div>
      <div class="form-group">
        <label for="exampleInputEmail2">身份证号</label>
        <input type="email" class="form-control" name="edit_time" id="exampleInputEmail2" value="{{.Search.edit_time}}" />
      </div>
      <div class="form-group">
        <label for="exampleInputEmail2">创建时间</label>
        <input type="email" class="form-control" name="edit_time" id="exampleInputEmail2" value="{{.Search.edit_time}}" />
      </div>
      <div class="form-group">
        <label for="exampleInputEmail2">编辑时间</label>
        <input type="email" class="form-control" name="edit_time" id="exampleInputEmail2" value="{{.Search.edit_time}}" />
      </div> 
      <div class="form-group">
        <label for="exampleInputEmail2">是否是测试账号</label>
        <input type="email" class="form-control" name="edit_time" id="exampleInputEmail2" value="{{.Search.edit_time}}" />
      </div>
      <div class="form-group">
        <label for="exampleInputEmail2">状态</label>
        <input type="email" class="form-control" name="edit_time" id="exampleInputEmail2" value="{{.Search.edit_time}}" />
      </div>
      <button type="button" id="search" class="btn btn-default">搜索</button>
      <button type="button" id="export" class="btn btn-default">导出</button>
    </form>
  </div>
  <!-- 按钮 -->
  <a href="//127.0.0.1:8080/novel/add" class="btn btn-primary btn-lg active" role="button">添加</a>
  <!-- 列表 -->
  <table class="table table-striped">
    <thead>
      <tr>
        <th scope="col">ID</th>
        <th scope="col">用户名</th>
        <th scope="col">真实姓名</th>
        <th scope="col">昵称</th>
        <th scope="col">性别</th>
        <th scope="col">年龄</th>
        <th scope="col">手机号</th>
        <th scope="col">身份证号</th>
        <th scope="col">创建时间</th>
        <th scope="col">编辑证号</th>
        <th scope="col">状态</th>
        <th scope="col">是否是测试账号</th>
        <th scope="col">操作</th>1
      </tr>
    </thead>
    <tbody>
      {{if .Page.HasContents }}
        {{range $index, $elem := .Page.List}}
        <tr>
          <th scope="row">{{$elem.Id}}</th>
          <td>{{$elem.UserName}}</td>
          <td>{{$elem.TrueName}}</td>
          <td>{{$elem.NickTime}}</td>
          <td>{{$elem.Sex}}</td>
          <td>{{$elem.Age}}</td>
          <td>{{$elem.Mobile}}</td>
          <td>{{$elem.Card}}</td>
          <td>{{$elem.CreateTime}}</td>
          <td>{{$elem.EditTime}}</td>
          <td>{{$elem.Status}}</td>
          <td>{{$elem.IsTest}}</td>
          <td></td>
        </tr>
        {{end}}
      {{else}}
        没有数据
      {{end}}
    </tbody>
    <!-- 页码 -->
  </table>
  <ul aria-label="Page navigation" id="page" style="float:right"></ul>
</div>
<script type="text/javascript">
//分页
  $(function () {
    {{if .Page.HasContents }}
      $("#page").bootstrapPaginator({
        currentPage: '{{.Page.PageNo}}',
        totalPages: '{{.Page.TotalPage}}',
        bootstrapMajorVersion: 3,
        size: "small",
        onPageClicked: function (e, originalEvent, type, page) {
            window.location.href = '{{.Page.PageUrl}}&PageNo=' + page
        }
      });
    {{end}}
    //提交表单
    $("#search").bind('click',function(){
      $("#search_data").attr('action','/employee');
      $("#search_data").submit();
    });
    $("#export").bind('click',function(){
      $("#search_data").attr('action','/employee/export');
      $("#search_data").submit();
    });
  });
</script>