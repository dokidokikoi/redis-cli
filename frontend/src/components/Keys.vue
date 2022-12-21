<template>
    <el-dialog 
    v-model="keyDiallogVisible" 
    title="创建键"
    width="60%">
        <el-form :model="form" label-width="120px">
            <el-form-item label="键的名称">
                <el-input placeholder="请输入键的名称" v-model="keyForm.key" />
            </el-form-item>
            <el-form-item label="键的类型">
                <el-select v-model="keyForm.type">
                    <el-option label="请选择数据类型" disabled></el-option>
                    <el-option :label="item" :value="item" v-for="item in keyTypes"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="createKey">创建</el-button>
                <el-button @click="keyDiallogVisible=false">取消</el-button>
            </el-form-item>
        </el-form>
    </el-dialog>
    <main>
        <el-form :inline="true" :model="form">
            <el-form-item>
                <el-input v-model="form.keyword" placeholder="请输入键的信息" @keyup.enter.stop="onSubmit" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit">查询</el-button>
            </el-form-item>
        </el-form>
        <el-button @click="keyDiallogVisible = true" style="width: 100%; margin-bottom: 12px;">新建</el-button>
        <div v-for="item in keys" @click="selectKeyKey(item)">
            <div v-if="item === selectKey">
                <div class="item key-select-item">
                    <div style="padding: 5px 12px;">{{ item }}</div>
                    <el-popconfirm title="确认删除?" @confirm="deleteKey(item)">
                        <template #reference>
                            <el-button text type="danger" @click.stop>删除</el-button>
                        </template>
                    </el-popconfirm>
                </div>
            </div>
            <div v-else>
                <div class="item key-item">
                    <div style="padding: 5px 12px;">{{ item }}</div>
                    <el-popconfirm title="确认删除?" @confirm="deleteKey(item)">
                        <template #reference>
                            <el-button text type="danger" @click.stop>删除</el-button>
                        </template>
                    </el-popconfirm>
                </div>
            </div>
        </div>
    </main>
</template>


<script setup>
import { reactive, ref, watch } from "vue"
import { DeleteKeyValue, KeyList, CreateKeyValue } from "../../wailsjs/go/main/App";
import { ElNotification } from 'element-plus'

let props = defineProps(['keyDB', 'keyIdentity'])
let emits = defineEmits(['emit-select-key'])

let form = reactive({
    keyword: ''
})

let keys = ref()

let selectKey = ref()

let keyTypes = ref([
    "string",
    "list",
    "hash",
    "set",
    "zset"
])
let keyForm = reactive({})
let keyDiallogVisible = ref(false)

watch(props, () => {
    getKeyList()
})

function onSubmit() {
    getKeyList()
}

function getKeyList() {
    let data = {
        conn_identity: props.keyIdentity,
        db: props.keyDB,
        keyword: form.keyword
    }
    KeyList(data).then(res => {
        if (res.code !== 200) {
            ElNotification({
                title: res.msg,
                type: "error"
            })
        }
        keys.value = res.data
    })
}

function selectKeyKey(key) {
    selectKey.value = key
    emits('emit-select-key', key)
}

function deleteKey(key) {
    DeleteKeyValue({conn_identity: props.keyIdentity, db: props.keyDB, key: key}).then(res =>{
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
        getKeyList()
    })
}

function createKey() {
    console.log({
        conn_identity: props.keyIdentity, 
        db: props.keyDB, 
        key: keyForm.key, 
        type: keyForm.type})
    CreateKeyValue({
        conn_identity: props.keyIdentity, 
        db: props.keyDB, 
        key: keyForm.key, 
        type: keyForm.type}).then(res => {
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
        keyDiallogVisible.value = false
        keyForm = undefined
        getKeyList()
    })
}
</script>


<style scoped>
.key-item {
    color: #409eff;
    background-color: #ecf5ff;
    margin-bottom: 5px;
}

.key-select-item {
    color: #67c23a;
    background-color: #f0f9eb;
    margin-bottom: 5px;
}
</style>