import axios from "axios";

function transJson2From(data) {
    // Do whatever you want to transform the data
    let ret = ''
    for (let it in data) {
        ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
    }
    return ret
}

const Api = {
    login(name,password,callback){
        axios({
            method: "POST",
            url: "/sso/login/exc",
            transformRequest: [transJson2From],
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