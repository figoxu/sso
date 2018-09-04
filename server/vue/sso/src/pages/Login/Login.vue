<template>
  <div className="login-page">
        <UserLogin />
      </div>
</template>

<script>
import UserLogin from './components/UserLogin';

export default {
  name: 'Login',
  components: {
    UserLogin,
  },
};
</script>
