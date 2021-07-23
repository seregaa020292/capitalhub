<template>
  <el-main>
    <h1 class="title text-center mb-3">Мои портфели</h1>
    <el-row :gutter="12">
      <el-col :md="6">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">Портфель 1</div>
          </template>
        </el-card>
      </el-col>
      <el-col :md="6">
        <el-card shadow="hover">
          <i class="el-icon-circle-plus" />
          Добавить еще портфель
        </el-card>
      </el-col>
    </el-row>
  </el-main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Portfolio',
})
</script>
