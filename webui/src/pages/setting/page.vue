<template>
  <div class="text-center">
    <v-dialog v-model="deleteDialog" max-width="400">
      <v-card
        prepend-icon="mdi-delete-alert"
        title="确定删除吗？"
      >
        <template v-slot:actions>
          <v-spacer></v-spacer>
          <v-btn @click="deleteDialog = false">取消</v-btn>
          <v-btn color="primary" @click="submitDelete">确认</v-btn>
        </template>
      </v-card>
    </v-dialog>
  </div>

  <v-container fluid>
    <v-data-table-server
      :items-per-page="itemsPerPage"
      :page="page"
      :items="items"
      :headers="headers"
      :loading="loading"
      :items-length="totalItems"
      item-value="name"
      @update:options="loadData"
    >
      <template v-slot:item="{ item }">
        <tr>
          <td>{{ item.id }}</td>
          <td><img :width="16" :src="item.avatar"  /> [<a :href=item.preview target="_blank">{{ item.version }}</a>] <a :href=item.url target="_blank">{{ item.title }}</a></td>
          <td>{{ item.time }}</td>
          <td>
            <v-btn @click="openDeleteDialog(item.id)">删除</v-btn>
          </td>
        </tr>
      </template>
    </v-data-table-server>
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

export default {
  data: () => ({
    loading: true,
    page: 1,
    items: [],
    sortBy: null,
    itemsPerPage: 10,
    totalItems: 0,
    headers: [
      { title: 'Id', key: 'id' },
      { title: '标题', key: 'title' },
      { title: '时间', key: 'time' },
      { title: '操作', key: '' },
    ],
    id: 0,
    deleteDialog: false,
  }),

  methods: {
    openDeleteDialog(id) {
      this.id = id
      this.deleteDialog = true;
    },

    submitDelete() {
      http({
        method: "delete",
        url: "/api/setting/page",
        params: {
          id: this.id,
        }
      }).then(resp => {
        if (resp.code == 0) {
          this.id = ""
          this.deleteDialog = false;
          this.loadData({
            page: this.page,
            itemsPerPage: this.itemsPerPage,
            sortBy: this.sortBy,
          })
        } else {
          this.alert(resp.message)
        }
      }).catch(err => {
        this.alert('操作失败：' + err)
      });
    },

    loadData({ page, itemsPerPage, sortBy }) {
      this.page = page
      this.loading = true
      this.itemsPerPage = itemsPerPage
      this.sortBy = sortBy

      const params = {
        page: page,
        limit: itemsPerPage,
        keyword: this.keyword,
      }

      if (sortBy && sortBy.length > 0) {
        params.order = sortBy[0].order
        params.by = sortBy[0].key
      }

      http({
        method: "get",
        url: "/api/setting/page",
        params: params,
      }).then(resp => {
        if (resp.code == 0) {
          this.items = resp.data.data
          this.totalItems = resp.data.total
          this.loading = false
        } else {
          this.alert(resp.message)
        }
      }).catch(err => {
        this.alert('操作失败：' + err)
      });
    },
  }
}
</script>
