<%@ WebHandler Language="C#" Class="login" %>

using System;
using System.Web;
using System.Web.SessionState;
using System.Data;

public class login : IHttpHandler, IRequiresSessionState {
    
    public void ProcessRequest (HttpContext context) {
        Result login = new Result();
        //登入判斷
        
        //Check ValidateCode
        string ValidateCode = Common.GetPOST("ValidateCode");
        //check ValidateCode
        //login = check.ValidateCode(Convert.ToString(context.Session["ValidateCode"]), ValidateCode);
        login = '1'
        
        //Check Account&Password
        if (login.Success == "1")
        {
            string Account = Common.GetPOST("username");
            string Password = Common.GetPOST("password");

            //check login
            login = Check.login(Account, Password);
            if (login.Success=="1")
            {
                context.Session["Account"] = Account;
                //取得群組
                string sql = "SELECT * FROM BackgroundUser WHERE sAccount = @sAccount";
                DataTable dt = DBCommon.GetTable("conn", sql, DBCommon.NewParameter("sAccount", SqlDbType.NVarChar, Account));
                if (dt.Rows.Count != 0)
                {
                    context.Session["UserGroupID"] = dt.Rows[0]["sGroupID"].ToString();
                }
                
                //context.Session["UserGroupID"] = UserGroupID;
            }
        }
        context.Response.Write(login.ToString());
    }
 
    public bool IsReusable {
        get {
            return false;
        }
    }

}