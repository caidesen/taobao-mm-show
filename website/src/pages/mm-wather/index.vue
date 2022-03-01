<script lang="ts" setup>
import {ref} from 'vue';
import {useQuasar} from 'quasar';
import {getNextPicture} from "../../service/app-service";

const $q = useQuasar();
const current = ref<MmPicture>();
const picList: MmPicture[]= [];
const loadMore = async () => {
  $q.loadingBar.start();
  try {
    const res = await getNextPicture();
    picList.push(...res)
    res.forEach(it => {
      const image = new Image();
      image.src = it.url
    })
  } finally {
    $q.loadingBar.stop();
  }
}
const onNext = async () => {
  if (!picList.length) {
    await loadMore()
  } else if (picList.length === 1) {
    loadMore()
  }
  current.value = picList.pop();
};
setTimeout(() => {
  onNext();
}, 100);
</script>
<template>
  <div>
    <q-img :src="current?.url" style="width: 100%; height: 100%">
      <div v-if="current?.text" class="absolute-bottom text-subtitle1 text-center">
        {{ current?.text }}
      </div>
    </q-img>
  </div>
  <q-page-sticky position="bottom-right" :offset="[18, 18]">
    <q-btn push fab icon="done_outline" @click="onNext" class="bg-info" color="#fff" />
  </q-page-sticky>
</template>

<style scoped></style>
