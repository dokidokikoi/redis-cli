<template>
    <main>
        <el-button :type="btnType" style="display: flex;" @click="dialogVisible=true">{{ title }}</el-button>
        <el-dialog 
            v-model="dialogVisible" 
            :title="title"
            width="60%">
            <el-form :model="form" label-width="120px">
                <el-form-item label="连接地址">
                    <el-input placeholder="请输入连接地址" v-model="form.addr" />
                </el-form-item>
                <el-form-item label="连接名称">
                    <el-input placeholder="请输入连接名称" v-model="form.name" />
                </el-form-item>
                <el-form-item label="连接端口">
                    <el-input placeholder="请输入连接端口" v-model="form.port" />
                </el-form-item>
                <el-form-item label="用户名">
                    <el-input placeholder="请输入用户名" v-model="form.username" />
                </el-form-item>
                <el-form-item label="密码">
                    <el-input placeholder="请输入连接端口" v-model="form.password" type="password" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" v-if="data === undefined" @click="createConnection">创建</el-button>
                    <el-button type="primary" v-else @click="editConnection">编辑</el-button>
                    <el-button @click="dialogVisible=false">取消</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </main>
</template>

<script setup>
import { reactive, ref } from "vue"
import { ConnectionCreate, ConnectionEdit } from "../../wailsjs/go/main/App"
import { ElNotification } from "element-plus"

let props = defineProps(['title', 'btnType', 'data'])
const emits = defineEmits(['emit-connection-list'])
const dialogVisible = ref(false)

let form = reactive({})

if (props.data !== undefined) {
    form = props.data
}

function createConnection() {
    ConnectionCreate(form).then(res => {
        if (res.code !== 200) {
            ElNotification({
                title: res.msg,
                type: "error"
            })
            return
        }
        ElNotification({
            title: res.msg,
            type: "success"
        })
        // 重新获取连接列表
        emits('emit-connection-list')
    })
}

function editConnection() {
    ConnectionEdit(form).then(res => {
        if (res.code !== 200) {
            ElNotification({
                title: res.msg,
                type: "error"
            })
            return
        }
        ElNotification({
            title: res.msg,
            type: "success"
        })
        // 重新获取连接列表
        emits('emit-connection-list')
    })
}
</script>

<style scoped>

</style>