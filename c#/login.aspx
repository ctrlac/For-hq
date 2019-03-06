<%@ Page Language="C#" AutoEventWireup="true" CodeFile="login.aspx.cs" Inherits="login" %>

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title>登入系統</title>
    <link rel="stylesheet" href="css/login.css" />
    <!--#include file="template/js.html"-->
    <script type="text/jscript" src="js/login/login.js"></script>
    <script type="text/jscript">
        function login() {
            $.ajax({
                type: "POST",
                url: "API/login.ashx",
                data: $('#form1').serialize(),
                dataType: "json",
                async: false,
                success: function (rs) {
                    if (rs.Success == "1") {
                        location.replace("main.aspx");
                    } else {
                        //fail_callback(rs.Msg);
                        alert(rs.Msg)
                    }
                },
                error: function () {
                    alert("ERROR!!!");
                }
            });
        }
        $(function () {
            $("form input").keypress(function (e) {
                if ((e.which && e.which == 13) || (e.keyCode && e.keyCode == 13)) {
                    login();
                    return false;
                }
            });
        });
    </script>
</head>
<body>
    <div class="container">
        <section id="content">
		<form id="form1" action="">
			<h1>登入系統</h1>
			<div>
				<input name="username" type="text" placeholder="帳號" required="" id="username" tabindex="1" />
			</div>
			<div>
				<input name="password" type="password" placeholder="密碼" required="" id="password" tabindex="2" />
			</div>
			<div>
                <img id="ValidateCode" src="ValidateCodeImage.ashx" />
                <a href="javascript:void(0);" onclick="ReSetCode();"><img src="images/login/btn_reserve.png" /></a>
                <input name="ValidateCode" type="text" placeholder="驗證碼" required="" tabindex="3" />
				<input type="button" value="登入" onclick="login();" tabindex="4" />
			</div>
		</form><!-- form -->
	    </section>
        <!-- content -->
    </div>
    <!-- container -->
</body>
</html>
