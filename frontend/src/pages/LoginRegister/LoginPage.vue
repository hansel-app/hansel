<template>
  <h1>Login</h1>
  <div class="container">
    <van-form @submit="onSubmit">
      <van-field
        v-model="username"
        name="Username"
        placeholder="Username"
        :rules="[{ required: true, message: 'Username is required' }]"
      />
      <van-field
        v-model="password"
        type="password"
        name="Password"
        placeholder="Password"
        :rules="[{ required: true, message: 'Password is required' }]"
      />
      <div style="margin: 16px;">
        <van-button round block type="primary" native-type="submit" width="80%">
          Login
        </van-button>
      </div>
    </van-form>
  </div>
</template>

<style scoped lang="less">
.container .van-field {
  width: 80%;
}
</style>
