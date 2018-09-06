<template>
    <div class="user-login">
        <div class="user-login-bg" :style="{'background-image':`url(${backgroundImage})`}"></div>
        <div class="content-wrapper">
            <h2 class="slogan">
                欢迎使用 <br/> 账户认证中心
            </h2>
            <div class="form-container">
                <h4 class="form-title">登录</h4>
                <el-form ref="form" :model="user" label-width="0">
                    <div class="form-items">
                        <el-row class="form-item">
                            <el-col>
                                <el-form-item prop="username" :rules="[ { required: true, message: '会员名/邮箱/手机号不能为空'}]">
                                    <div class="form-line">
                                        <i class="el-icon-edit-outline input-icon"></i>
                                        <el-input placeholder="会员名/邮箱/手机号" v-model="user.username"></el-input>
                                    </div>
                                </el-form-item>
                            </el-col>
                        </el-row>
                        <el-row class="form-item" style='margin-top:10px;'>
                            <el-col>
                                <el-form-item prop="password" :rules="[ { required: true, message: '密码不能为空'}]">
                                    <div class="form-line">
                                        <i class="el-icon-service input-icon"></i>
                                        <el-input type="password" placeholder="密码" v-model="user.password"></el-input>
                                    </div>
                                </el-form-item>
                            </el-col>
                        </el-row>
                        <el-row class="form-item" style='margin-top:10px;'>
                            <el-button type="primary" class="submit-btn" size="small" @click="submitBtn">
                                登 录
                            </el-button>
                        </el-row>
                    </div>
                </el-form>
            </div>
        </div>
    </div>
</template>

<script>
    import Api from "../../../../api/api";
    import BasicContainer from '@vue-materials/basic-container';
    const backgroundImage = '/images/login_bg.png';
    export default {
        components: {BasicContainer},
        name: 'UserLogin',

        data() {
            return {
                backgroundImage: backgroundImage,
                user: {
                    username: '',
                    password: '',
                },
            };
        },

        created() {
        },

        methods: {
            submitBtn() {
                var that = this;
                this.$refs['form'].validate((valid) => {
                    if (valid) {
                        Api.login(that.user.username, that.user.password, function (data) {
                            console.log("调用成功", data)
                        });
                        this.$message({
                            message: '登录成功',
                            type: 'success',
                        });
                    }
                });
            },
        },
    };
</script>

<style lang="scss" scoped>
    @import './UserLogin.scss';
</style>