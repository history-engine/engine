<template>
  <div class="text-center">
    <v-dialog v-model="addDialog" max-width="600" persistent>
      <v-card prepend-icon="mdi-account" title="添加host">
        <v-card-text>
          <v-row dense>
            <v-col cols="12" sm="12">
              <v-select
                :items="types"
                label="类型"
                v-model="type"
                item-title="state"
                item-value="abbr"
                return-object
                required
              ></v-select>
            </v-col>

            <v-col cols="12" sm="12">
              <v-text-field
                hint="支持完整域名a.b.cn、通配符*.b.cn、正则regexp:.*(\.|)(a|b|c|d)\.cn"
                label="host"
                v-model="host"
                required
              ></v-text-field>
            </v-col>

          </v-row>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn
            text="取消"
            variant="plain"
            @click="addDialog = false"
          ></v-btn>

          <v-btn
            color="primary"
            text="提交"
            variant="tonal"
            @click="submitAdd"
          ></v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="editDialog" max-width="600" persistent>
      <v-card prepend-icon="mdi-account" title="修改host">
        <v-card-text>
          <v-row dense>
            <v-col cols="12" sm="12">
              <v-text-field
                label="Id*"
                v-model="id"
                readonly
              ></v-text-field>
            </v-col>

            <v-col cols="12" sm="12">
              <v-select
                :items="types"
                label="类型"
                v-model="type"
                :hint="`${type.state}, ${type.abbr}`"
                item-title="state"
                item-value="abbr"
                return-object
                required
              ></v-select>
            </v-col>

            <v-col cols="12" sm="12">
              <v-text-field
                hint="支持完整域名a.b.cn、通配符*.b.cn、正则regexp:.*(\.|)(a|b|c|d)\.cn"
                label="host"
                v-model="host"
                required
              ></v-text-field>
            </v-col>

          </v-row>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn
            text="取消"
            variant="plain"
            @click="editDialog = false"
          ></v-btn>

          <v-btn
            color="primary"
            text="提交"
            variant="tonal"
            @click="submitEdit"
          ></v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

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
    <v-row align="center" justify="end">
      <v-col cols="8">
        <v-text-field
          v-model="keyword"
          prepend-inner-icon="mdi-magnify"
          variant="solo-filled"
          flat
          hide-details
          @keydown.enter="search"
        ></v-text-field>
      </v-col>

      <v-col cols="4">
        <v-btn size="large" variant="tonal" @click="search" class="mr-2">搜索</v-btn>
        <v-btn size="large" variant="tonal" @click="openAddDialog">添加</v-btn>
      </v-col>
    </v-row>

    <v-data-table-server
      :items-per-page="itemsPerPage"
      :page="page"
      :items="items"
      :headers="headers"
      :loading="loading"
      :items-length="totalItems"
      item-value="name"
      @update:options="loadItems"
    >
      <template v-slot:item="{ item }">
        <tr>
          <td>{{ item.id }}</td>
          <td>{{ item.host }}</td>
          <td>{{ item.type == 1 ? '包括' : '忽略' }}</td>
          <td>{{ item.created_at }}</td>
          <td>
            <v-btn @click="openDeleteDialog(item.id)" class="mr-2">删除</v-btn>
            <v-btn @click="openEditDialog(item.id, item.type, item.host)">修改</v-btn>
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
import http from "@/services/http"

export default {
  data: () => ({
    page: 1,
    items: [],
    itemsPerPage: 10,
    totalItems: 0,
    sortBy: null,
    headers: [
      { title: 'Id', key: 'id' },
      { title: '域名', key: 'host' },
      { title: '类型', key: 'type', value: v => { return v == 1 ? '包括' : '忽略'} },
      { title: '更新时间', key: 'updated_at' },
      { title: '操作', key: '' },
    ],
    keyword: '',
    loading: true,
    types: [
      { state: '包括', abbr: 1 },
      { state: '忽略', abbr: 2 },
    ],
    id: 0,
    type: {},
    host: '',
    addDialog: false,
    deleteDialog: false,
    editDialog: false,
  }),

  methods: {
    openAddDialog() {
      this.type = { state: '忽略', abbr: 2 }
      this.host = ""
      this.addDialog = true
    },

    openDeleteDialog(id) {
      this.id = id
      this.deleteDialog = true;
    },

    openEditDialog(id, type, host) {
      this.id = id
      this.type = { state: type == 1 ?  '包括' : '忽略', abbr: type }
      this.host = host
      this.editDialog = true;
    },

    submitAdd() {
      http({
        method: "put",
        url: "/api/setting/host",
        data: {
          type: this.type.abbr,
          host: this.host,
        }
      }).then(resp => {
        if (resp.code == 0) {
          this.search()
        } else {
          this.alert(resp.message)
        }
      }).catch(err => {
        this.alert('操作失败：' + err)
      });

      this.host = ''
      this.addDialog = false
    },

    submitDelete() {
      http({
        method: "delete",
        url: "/api/setting/host",
        params: {
          id: this.id,
        }
      }).then(resp => {
        if (resp.code == 0) {
          this.search()
        } else {
          this.alert(resp.message)
        }
      }).catch(err => {
        this.alert('操作失败：' + err)
      });

      this.id = 0
      this.host = ''
      this.deleteDialog = false;
    },

    submitEdit() {
      http({
        method: "post",
        url: "/api/setting/host",
        data: {
          id: this.id,
          type: this.type.abbr,
          host: this.host,
        }
      }).then(resp => {
        if (resp.code == 0) {
          this.search()
          this.id = 0
          this.host = ''
          this.editDialog = false;
        } else {
          this.alert(resp.message)
        }
      }).catch(err => {
        this.alert('操作失败：' + err)
      });
    },

    search() {
      this.loadItems({
        page:this.page,
        itemsPerPage:this.itemsPerPage,
        sortBy:this.sortBy,
      })
    },

    loadItems ({ page, itemsPerPage, sortBy }) {
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
        url: "/api/setting/host",
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

    alert(text) {
      this.snackbar = true
      this.text = text
    },
  },
}
</script>
