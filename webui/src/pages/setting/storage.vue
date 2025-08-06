<template>
  <div class="text-center">
    <v-snackbar v-model="snackbar" :timeout="timeout" location="top">
      {{ text }}

      <template v-slot:actions>
        <v-btn
          color="blue"
          variant="text"
          @click="snackbar = false"
        >
          关闭
        </v-btn>
      </template>
    </v-snackbar>
  </div>

  <v-container fluid>
    <v-form validate-on="submit" rounded="lg" @submit.prevent>
      <div class="text-subtitle-1 text-medium-emphasis">最大版本数</div>
      <v-number-input
        class="mb-4"
        hint="超过设置的版本数时，将自动删除最早的版本。"
        :rules="[rules.required]"
        required
        :min="1"
        variant="outlined"
        v-model="setting.max_version"
      ></v-number-input>

      <div class="text-subtitle-1 text-medium-emphasis">最小保存间隔（秒）</div>
      <v-number-input
        class="mb-4"
        hint="此间隔时间内，重复访问一个网页，不再保存。"
        :rules="[rules.required]"
        required
        type="number"
        variant="outlined"
        v-model="setting.min_version_interval"
      ></v-number-input>

      <div class="text-subtitle-1 text-medium-emphasis">最小HTML文件大小</div>
      <v-number-input
        class="mb-4"
        hint="小于此大小的HTML文件不会被保存，1MB=1024*1024=1,048,576字节"
        :rules="[rules.required]"
        required
        type="number"
        variant="outlined"
        v-model="setting.min_size"
      ></v-number-input>

      <div class="text-subtitle-1 text-medium-emphasis">最大HTML文件大小</div>
      <v-number-input
        hint="大于此大小的HTML文件不会被保存，1MB=1024*1024=1,048,576字节"
        :rules="[rules.required]"
        required
        type="number"
        variant="outlined"
        v-model="setting.max_size"
      ></v-number-input>

      <v-btn variant="outlined" @click="submit">保存</v-btn>
    </v-form>
  </v-container>
</template>

<route>
{
meta: {
layout: "setting"
}
}
</route>

<script>
import http from "@/services/http";
import {set} from "core-js/internals/task";

export default {
  data: () => ({
    snackbar: false,
    text: "",
    timeout: 3000,
    rules: {
      required: value => !!value || 'Field is required',
    },
    setting: {
      max_version: 0,
      min_version_interval: 0,
      min_size: 0,
      max_size: 0,
    }
  }),

  methods: {
    set,
    submit() {
      http({
        method: "post",
        url: "/api/setting/storage",
        data: {
          max_version: this.setting.max_version,
          min_version_interval: this.setting.min_version_interval,
          min_size: this.setting.min_size,
          max_size: this.setting.max_size,
        }
      }).then(resp => {
        if (resp.code == 0) {
          this.alert(resp.message)
        } else {
          this.alert(resp.message)
        }
      }).catch(err => {
        this.alert('操作失败：' + err)
      });
    },

    loadFormData() {
      http({
        method: "get",
        url: "/api/setting/storage",
      }).then(resp => {
        if (resp.code == 0) {
          this.setting = resp.data
        } else {
          this.alert(resp.message)
        }
      }).catch(err => {
        this.alert('操作失败：' + err)
      });
    },

    alert(text) {
      this.snackbar = true
      this.text = text
    },
  },

  created() {
    this.loadFormData();
  }
}
</script>
