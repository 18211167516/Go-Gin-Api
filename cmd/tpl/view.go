package tpl

func ViewTemplate() string {
	return `
{{define "<!--{.Default}-->.html"}}
{{template "common/header" .}}
<link href="/public/static/js/bootstrap-table/bootstrap-table.min.css" rel="stylesheet">
</head>
<body>
<div class="container-fluid p-t-15">
	<div class="row">
	<div class="col-lg-12">
		<div class="card">
		<header class="card-header"><div class="card-title"><!--{.Desc}-->管理</div></header>
		<div class="card-body">
			<div id="toolbar" class="toolbar-btn-action">
			<form method="GET" action="/" id="Filter" class="form-inline">
				<div class="form-group">
				<button id="btn_add" type="button" class="btn btn-primary m-r-5 {{call .addUrl.Allow}}" data-toggle="modal" data-target="#add">
					<span class="mdi mdi-plus" aria-hidden="true"></span>新增
				</button>
				</div>
				<div class="form-group">
				<label class="" for="name">名称</label>
				<div class="input-group">
					<div class="input-group-addon"><i class="glyphicon glyphicon-search"></i></div>
					<input type="text" class="form-control input-sm" name="name" placeholder="名称">
				</div>
				</div>
				<div class="form-group">
				<label for="type" class="">状态</label>
				<div class="col-sm-8">
					<select name="status" class="form-control" id="status">
					<option value="">请选择</option>
					<option value="1">启用</option>
					<option value="2">其他状态</option>
					</select>
				</div>
				</div>
				<div class="form-group">
				<button type="submit" class="btn btn-primary btn-sm">提交</button>
				</div>
			</form>
			</div>
			<table id="tb_departments"></table>
			
		</div>
		</div>
	</div>
	</div>
	<div class="modal fade" tabindex="-1" role="dialog" id="add">
	<div class="modal-dialog modal-dialog-centered" role="document">
		<div class="modal-content">
		<div class="modal-header">
			<h6 class="modal-title" id="exampleModalCenterTitle">创建<!--{.Desc}--></h6>
			<button type="button" class="close" data-dismiss="modal" aria-label="Close">
			<span aria-hidden="true">&times;</span>
			</button>
		</div>
		<form class="form-horizontal needs-validation ajax-form" name="addForm" id="addForm" novalidate action="{{.addUrl.Url}}">
			<div class="modal-body">
			<div class="form-group row">
				<label for="name" class="control-label col-sm-3">名称</label>
				<div class="col-sm-8">
				<input type="text" class="form-control" name="name" required>
				<div class="invalid-feedback">
					名称不能为空
				</div>
				</div>
			</div>
			</div>
			<div class="modal-footer">
			<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
			<button type="submit" class="btn btn-primary">确定</button>
			</div>
		</form>
		</div><!-- /.modal-content -->
	</div><!-- /.modal-dialog -->
	</div>

	<div class="modal fade" tabindex="-1" role="dialog" id="edit">
	<div class="modal-dialog modal-dialog-centered" role="document">
		<div class="modal-content">
		<div class="modal-header">
			<h6 class="modal-title" id="exampleModalCenterTitle">编辑<!--{.Desc}--></h6>
			<button type="button" class="close" data-dismiss="modal" aria-label="Close">
			<span aria-hidden="true">&times;</span>
			</button>
		</div>
		<form class="form-horizontal needs-validation ajax-form" name="editForm" id="editForm" novalidate action="">
			<div class="modal-body">
			<div class="form-group row">
				<label for="role_name" class="control-label col-sm-3">名称</label>
				<div class="col-sm-8">
				<input type="text" class="form-control" name="name" required>
				<div class="invalid-feedback">
					名称不能为空
				</div>
				</div>
			</div>
			<div class="form-group row">
				<label for="status" class="control-label col-sm-3">状态</label>
				<div class="col-sm-8">
				<select name="status" class="form-control" id="status">
					<option value="1">启用</option>
					<option value="0">禁用</option>
				</select>
			</div>
			</div>
			</div>
			<div class="modal-footer">
			<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
			<button type="submit" class="btn btn-primary">确定</button>
			</div>
		</form>
		</div><!-- /.modal-content -->
	</div><!-- /.modal-dialog -->
	</div>

	
</div>

{{template "common/footer"}}
<script src="https://cdn.jsdelivr.net/npm/tableexport.jquery.plugin@1.10.21/tableExport.min.js"></script>
<script type="text/javascript" src="/public/static/js/bootstrap-table/extensions/export/bootstrap-table-export.min.js"></script>
<script type="text/javascript">
$table = $('#tb_departments')
$table.bootstrapTable({
	classes: 'table table-bordered table-hover table-striped',
	url: {{.dataUrl}},
	method: {{.dataMethod}},
	contentType:"application/x-www-form-urlencoded",
	dataType : 'json',        // 因为本示例中是跨域的调用,所以涉及到ajax都采用jsonp,
	responseHandler: function(res) {//定义返回格式
		return {
			"total": res.data.total,//总页数
			"rows": res.data.lists   //数据
			};
	},
	uniqueId: 'id',
	idField: 'id',             // 每行的唯一标识字段
	toolbar: '#toolbar',       // 工具按钮容器
	//clickToSelect: true,     // 是否启用点击选中行
	//showColumns: true,         // 是否显示所有的列
	showRefresh: true,         // 是否显示刷新按钮
	//showToggle: true,        // 是否显示详细视图和列表视图的切换按钮(clickToSelect同时设置为true时点击会报错)
	pagination: true,                    // 是否显示分页
	sortOrder: "asc",                    // 排序方式
	queryParams: function(params) {
		var filter = $('#Filter').serializeArray();
		var temp = {
			limit: params.limit,         // 每页数据量
			page: (params.offset / params.limit) + 1,
			sort: params.sort,           // 排序的列名
			sortOrder: params.order,     // 排序方式'asc' 'desc'
		};
		for (var i in filter) {
			temp[filter[i].name] =filter[i].value
		}
		return temp;
	},                                   // 传递参数
	sidePagination: "server",            // 分页方式：client客户端分页，server服务端分页
	pageNumber: 1,                       // 初始化加载第一页，默认第一页
	pageSize: 10,                        // 每页的记录行数
	pageList: [10, 25, 50, 100],         // 可供选择的每页的行数
	//search: true,                      // 是否显示表格搜索，此搜索是客户端搜索
	showExport: true,        // 是否显示导出按钮, 导出功能需要导出插件支持(tableexport.min.js)
	exportDataType: "all", // 导出数据类型, 'basic':当前页, 'all':所有数据, 'selected':选中的数据
	exportOptions: {//导出设置
		fileName: 'Tablexxx',//下载文件名称
		ignoreColumn: [6] //忽略某一列导出
	},
	columns: [{
		field: 'example',
		checkbox: true    // 是否显示复选框
	}, {
		field: 'id',
		title: 'ID',
		sortable: true    // 是否排序
	}, {
		field: 'name',
		title: '名称',
	},
	{
		field: 'created_at',
		title: '创建时间'
	}, {
		field: 'updated_at',
		title: '修改时间'
	}, {
		field: 'status',
		title: '状态',
		formatter:function(value, row, index){ 
			var value="";
			if (row.status == '0') {
				value = '<span class="badge badge-danger">禁用</span>';
			} else if(row.status == '1') {
				value = '<span class="badge badge-success">正常</span>';
			}else {
				value = row.status ;
			}
			return value;
		}
	}, {
		field: 'operate',
		title: '操作',
		formatter: function(value,row,index){
			return btnGroup(value,row, index)
		},  // 自定义方法
		events: {
			'click .edit-btn': function (event, value, row, index) {
				editUser(row);
			},
		}
	}],
	onLoadSuccess: function(data){
		$("[data-toggle='tooltip']").tooltip();
	}
});

$('#Filter').submit(function(e){
	e.preventDefault();
	$table.bootstrapTable('refresh');
});

// 操作按钮
function btnGroup (value, row, index)
{
	delUrl = "{{.delUrl.Url}}";
	delUrl = delUrl.replace(/(:\w+)/,row.id)
	console.log(delUrl)
	let html = '';

	html +=
		'<a href="#edit" class="btn btn-xs btn-default m-r-5 edit-btn {{call .editUrl.Allow}}" title="编辑" data-toggle="tooltip"><i class="mdi mdi-pencil"></i></a>' +
		'<a href="'+delUrl+'" class="btn btn-xs btn-default del-btn ajax-post confirm {{call .delUrl.Allow}}" data-tips="确认删除吗" title="删除"><i class="mdi mdi-window-close"></i></a>';
	
	return html;
}

// 操作方法 - 编辑
function editUser(row)
{
	url = "{{.editUrl.Url}}";
	url = url.replace(/(:\w+)/,row.id)
	console.log(url)
	name = row.name;
	status = row.status;
	$('#edit').find('#editForm').attr('action',url);
	$('#edit').find('input[name="name"]').val(name);
	$('#edit').find('select[name="status"]').val(status);
	$('#edit').modal({show:true})
}

</script>
</body>
</html>
{{end}}
`
}
