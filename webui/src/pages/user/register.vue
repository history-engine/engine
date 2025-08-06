<template>
  <div class="auth-wrapper d-flex align-center justify-center pa-4">
    <VCard class="auth-card pa-4 pt-7" max-width="448">
      <VCardText class="pt-2">
        <h5 class="text-h5 mb-1 font-weight-bold">
          欢迎来到History Engine！ 👋🏻
        </h5>
      </VCardText>

      <VCardText>
        <v-form ref="form" lazy-validation>
          <VRow>
            <!-- Username -->
            <VCol cols="12">
              <VTextField
                v-model="username"
                autofocus
                label="用户名"
                placeholder="需要保证唯一性"
              />
            </VCol>
            <!-- email -->
            <VCol cols="12">
              <VTextField
                v-model="email"
                label="邮箱"
                placeholder="尽量使用Hotmail、Gmail等常用邮箱"
                type="email"
              />
            </VCol>

            <!-- password -->
            <VCol cols="12">
              <VTextField
                v-model="password"
                label="密码"
                placeholder="8位以上数字、符号和字母的组合"
                :type="isPasswordVisible ? 'text' : 'password'"
                :append-inner-icon="isPasswordVisible ? 'mdi-eye-off-outline' : 'mdi-eye-outline'"
                @click:append-inner="isPasswordVisible = !isPasswordVisible"
              />
              <!--<div class="d-flex align-center mt-1 mb-4">
                <VCheckbox
                  id="privacy-policy"
                  v-model="privacyPolicies"
                  inline
                />
                <VLabel
                  for="privacy-policy"
                  style="opacity: 1;"
                >
                  <span class="me-1">我同意并愿意遵守</span>
                  <a
                    href="javascript:void(0)"
                    class="text-primary me-1"
                  >用户协议</a>
                  <span class="me-1">与</span>
                  <a
                    href="javascript:void(0)"
                    class="text-primary"
                  >隐私政策</a>
                </VLabel>
              </div>-->

              <v-btn @click="submit" block>注册</v-btn>
            </VCol>

            <!-- login instead -->
            <VCol cols="12" class="text-center text-base">
              <span>已有账号？</span>
              <RouterLink class="text-primary ms-2" to="/user/login">立即登录</RouterLink>
            </VCol>
          </VRow>
        </v-form>
      </VCardText>
    </VCard>
  </div>
</template>

<script>
import { useAppStore } from "@/stores/app";
import http from "@/services/http"
export default {
  setup() {
    const store = useAppStore();
    return {store}
  },

  data: () => ({
    username: '',
    email: '',
    password: '',
    privacyPolicies: false,
    isPasswordVisible: false,
  }),
  methods: {
    async submit() {
      if (!this.username || !this.email || !this.password) {
        alert("不能为空");
        return
      }
      http({
        method: 'post',
        url: "/api/user/register",
        data: {
          username: this.username,
          email: this.email,
          password: this.password,
        }
      }).then(res => {
        if (res.code == 0) {
          this.store.login(res.data.user);
          this.$router.push('/');
        } else {
          alert("注册失败：" + res.message)
        }
      }).catch(err => {
        alert("注册失败：" + err)
      });
    },
  },
};
</script>
