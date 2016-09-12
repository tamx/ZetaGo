<!DOCTYPE html>

<html>
  	<head>
    	<title>ZetaGo</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	
	</head>
  	
  	<body>
		<table id="board" border="1" bordercolor="BLACK">
			<tr>
				<td>{{.B11}}</td>
				<td>{{.B12}}</td>
				<td>{{.B13}}</td>
				<td>{{.B14}}</td>
				<td>{{.B15}}</td>
				<td>{{.B16}}</td>
				<td>{{.B17}}</td>
				<td>{{.B18}}</td>
			</tr>
			<tr>
				<td>{{.B21}}</td>
				<td>{{.B22}}</td>
				<td>{{.B23}}</td>
				<td>{{.B24}}</td>
				<td>{{.B25}}</td>
				<td>{{.B26}}</td>
				<td>{{.B27}}</td>
				<td>{{.B28}}</td>
			</tr>
			<tr>
				<td>{{.B31}}</td>
				<td>{{.B32}}</td>
				<td>{{.B33}}</td>
				<td>{{.B34}}</td>
				<td>{{.B35}}</td>
				<td>{{.B36}}</td>
				<td>{{.B37}}</td>
				<td>{{.B38}}</td>
			</tr>
			<tr>
				<td>{{.B41}}</td>
				<td>{{.B42}}</td>
				<td>{{.B43}}</td>
				<td>{{.B44}}</td>
				<td>{{.B45}}</td>
				<td>{{.B46}}</td>
				<td>{{.B47}}</td>
				<td>{{.B48}}</td>
			</tr>
			<tr>
				<td>{{.B51}}</td>
				<td>{{.B52}}</td>
				<td>{{.B53}}</td>
				<td>{{.B54}}</td>
				<td>{{.B55}}</td>
				<td>{{.B56}}</td>
				<td>{{.B57}}</td>
				<td>{{.B58}}</td>
			</tr>
			<tr>
				<td>{{.B61}}</td>
				<td>{{.B62}}</td>
				<td>{{.B63}}</td>
				<td>{{.B64}}</td>
				<td>{{.B65}}</td>
				<td>{{.B66}}</td>
				<td>{{.B67}}</td>
				<td>{{.B68}}</td>
			</tr>
			<tr>
				<td>{{.B71}}</td>
				<td>{{.B72}}</td>
				<td>{{.B73}}</td>
				<td>{{.B74}}</td>
				<td>{{.B75}}</td>
				<td>{{.B76}}</td>
				<td>{{.B77}}</td>
				<td>{{.B78}}</td>
			</tr>
			<tr>
				<td>{{.B81}}</td>
				<td>{{.B82}}</td>
				<td>{{.B83}}</td>
				<td>{{.B84}}</td>
				<td>{{.B85}}</td>
				<td>{{.B86}}</td>
				<td>{{.B87}}</td>
				<td>{{.B88}}</td>
			</tr>
		</table>
		<a href="/?reset=true">Reset</a>
		<hr/>
		{{.log}}
		<script type="text/javascript">
		<!--
		var table = document.getElementById('board');
		for (var y = 1; y <= 8; y++) {
			for (var x = 1; x <= 8; x++) {
				var cell = table.rows[y-1].cells[x-1];
				var data = cell.firstChild.data;
				var elem = document.createElement("a");
				elem.href = "/?p=" + y + x + "1";
				var str = document.createTextNode(data);
				elem.appendChild(str);
				cell.removeChild(cell.childNodes.item(0));
				cell.appendChild(elem);
			}
		}
		-->
		</script>
	</body>
</html>
