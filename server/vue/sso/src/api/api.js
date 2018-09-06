import axios from "axios";

const Api = {
    login(name,password,callback){
        axios({
            method: "POST",
            url: "/sso/login/exc",
            data: {
                username: name,
                password: password
            }
        }).then(function (res) {
            console.log(">>>>");
            console.log(res);
            // Message.success("登录成功");
            // res.data.Authorization = "Basic " + util.base64Encode(res.data.name + ":" + res.data.token);
            // Cookies.set("userInfo", res.data);
            if (callback) callback(res.data);
        }).catch(function (error) {
            // if (error.response) pushErrorNotice(error.response.data);
        });
    }
};




export default Api;