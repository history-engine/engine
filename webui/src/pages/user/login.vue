<template>
  <div class="auth-wrapper d-flex align-center justify-center pa-4">
    <VCard
      class="auth-card pa-4 pt-7"
      max-width="448"
    >
      <v-card-text class="pt-2">
        <h5 class="text-h5 mb-1 font-weight-bold">
          欢迎回到History Engine！ 👋🏻
        </h5>
      </v-card-text>

      <v-card-text>
        <v-form validate-on="submit lazy" @submit.prevent="submit">
          <v-row>
            <!-- email -->
            <v-col cols="12">
              <VTextField
                v-model="username"
                autofocus
                placeholder=""
                label="用户名"
                type="username"
              />
            </v-col>

            <!-- password -->
            <v-col cols="12">
              <v-text-field
                v-model="password"
                label="密码"
                placeholder=""
                :type="isPasswordVisible ? 'text' : 'password'"
                :append-inner-icon="isPasswordVisible ? 'mdi-eye-off-outline' : 'mdi-eye-outline'"
                @click:append-inner="isPasswordVisible = !isPasswordVisible"
              />

              <!-- remember me checkbox -->
              <div class="d-flex align-center justify-space-between flex-wrap mt-1 mb-4">
                <!--<v-checkbox
                  v-model="remember"
                  label="记住我"
                />-->

                <RouterLink
                  class="text-primary ms-2 mb-1"
                  to="/user/rest-password"
                >
                  忘记密码？
                </RouterLink>
              </div>

              <!-- login button -->
              <VBtn
                block
                type="submit"
              >
                登录
              </VBtn>
            </v-col>

            <!-- create account -->
            <VCol
              cols="12"
              class="text-center text-base"
            >
              <span>没有账号？</span>
              <RouterLink
                class="text-primary ms-2"
                to="/user/register"
              >
                立即注册
              </RouterLink>
            </VCol>
          </v-row>
        </v-form>
      </v-card-text>
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
    username: "",
    password: "",
    remember: false,
    isPasswordVisible: false,
  }),

  methods: {
    async submit() {
      http({
        method: 'post',
        url: "/api/user/login",
        data: {
          username: this.username,
          password: this.password,
        },
      }).then(res => {
        if (res.code == 0) {
          this.store.login(res.data.user);
          this.$router.push('/') // todo 回到之前的页面
        } else {
          alert(res.message)
        }
      }).catch(err => {
        alert('登录失败：' + err)
      });
    },
  },
};
</script>
