<template>
  <v-container class="">
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
          <td>{{ item.suffix }}</td>
          <td>{{ item.type == 1 ? '包括' : '忽略' }}</td>
          <td>{{ item.created_at }}</td>
          <td>
            <v-btn @click="openDeleteDialog(item.id)" class="mr-2">删除</v-btn>
            <v-btn @click="openEditDialog(item.id, item.type, item.suffix)">修改</v-btn>
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
      { title: '后缀', key: 'suffix' },
      { title: '类型', key: 'type', value: v => { return v == 1 ? '包括' : '忽略'} },
      { title: '更新时间', key: 'updated_at' },
      { title: '操作', key: '' },
    ],
  }),

  methods: {
    openEditDialog() {
      // todo
    },

    openDeleteDialog() {
      //  todo
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
        url: "/api/setting/filetype",
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
    }
  }
}
</script>
