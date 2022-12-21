<template>
    <main>
        <div class="demo-collapse">
            <el-collapse accordion>
                <el-collapse-item :name="item.identity" v-for="item in list" @click="getInfo(item.identity)">
                    <template #title>
                        <div class="item">
                            <div>
                                {{ item.name }}
                            </div>
                            <div style="display: flex; align-items: center;">
                                <ConnectionManage @click.stop title="编辑" btn-type="text" :data="item" />
                                <el-popconfirm title="确认删除?" @confirm="connectionDelete(item.identity)">
                                    <template #reference>
                                        <el-button link type="danger" @click.stop>删除</el-button>
                                    </template>
                                </el-popconfirm>
                            </div>
                        </div>
                    </template>
                    <div v-for="db in infoDBList" @click="selectDB(db.key, item.identity)">
                        <div v-if="selectDBKey !== db.key" class="my-item">{{ db.key }} ({{ db.number }})</div>
                        <div v-else class="my-select-item">{{ db.key }} ({{ db.number }})</div>
                    </div>
                </el-collapse-item>
            </el-collapse>
        </div>
    </main>
</template>


<script setup>
import { ref, watch } from "vue"
import { ConnectionList, ConnectionDelete, DBList } from '../../wailsjs/go/main/App'
import { ElNotification } from "element-plus"
import ConnectionManage from "./ConnectionManage.vue";

const props = defineProps(['flush'])
const emits = defineEmits(['emit-select-db'])
let list = ref()

let infoDBList = ref()

let selectDBKey = ref()

watch(props, (newFlush) => {
    connectionList()
})

function connectionList() {
    ConnectionList().then(res => {
        if (res.code !== 200) {
            ElNotification({
                title: res.msg,
                type: "error"
            })
        }
        list.value = res.data
    })
}

connectionList()

function connectionDelete(identity) {
    ConnectionDelete(identity).then(res =>{
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
        connectionList()
    })
}

function getInfo(identity) {
    DBList(identity).then(res => {
        if (res.code !== 200) {
            ElNotification({
                title: res.msg,
                type: "error"
            })
        }
        infoDBList.value = res.data
    })
}

function selectDB(db, identity) {
    selectDBKey = db
    emits("emit-select-db", Number(db.substring(2)), identity)
}

</script>


<style scoped>
</style>