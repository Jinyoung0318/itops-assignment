<template>
  <div>
    <h2>{{ isEdit ? '이슈 수정' : '이슈 생성' }}</h2>
    <form @submit.prevent="handleSubmit">
      <div>
        <label>제목:</label>
        <input v-model="form.title" required />
      </div>
      <div>
        <label>설명:</label>
        <textarea v-model="form.description" required></textarea>
      </div>
      <div>
        <label>담당자:</label>
        <select v-model="form.userId" :disabled="isCompletedOrCancelled">
          <option :value="null">없음</option>
          <option v-for="user in users" :key="user.id" :value="user.id">{{ user.name }}</option>
        </select>
      </div>
      <div>
        <label>상태:</label>
        <select v-model="form.status" :disabled="!form.userId || isCompletedOrCancelled">
          <option value="PENDING">PENDING</option>
          <option value="IN_PROGRESS">IN_PROGRESS</option>
          <option value="COMPLETED">COMPLETED</option>
          <option value="CANCELLED">CANCELLED</option>
        </select>
      </div>
      <button type="submit">저장</button>
      <button type="button" @click="goBack">목록으로</button>
    </form>
  </div>
</template>

<script>
import { ref, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import axios from 'axios';

export default {
  name: 'IssueForm',
  setup() {
    const router = useRouter();
    const route = useRoute();
    const isEdit = computed(() => !!route.params.id);
    const users = ref([
      { id: 1, name: '김개발' },
      { id: 2, name: '이디자인' },
      { id: 3, name: '박기획' },
    ]);
    const form = ref({
      title: '',
      description: '',
      status: 'PENDING',
      userId: null,
    });

    const isCompletedOrCancelled = computed(() => {
      return form.value.status === 'COMPLETED' || form.value.status === 'CANCELLED';
    });

    const goBack = () => {
      router.push('/issues');
    };

    const loadIssue = async (id) => {
      const res = await axios.get(`http://localhost:8080/issue/${id}`);
      const issue = res.data;
      form.value = {
        title: issue.title,
        description: issue.description,
        status: issue.status,
        userId: issue.user?.id ?? null,
      };
    };

    const handleSubmit = async () => {
      if (isEdit.value) {
        await axios.patch(`http://localhost:8080/issue/${route.params.id}`, form.value);
      } else {
        await axios.post('http://localhost:8080/issue', form.value);
      }
      goBack();
    };

    if (isEdit.value) {
      loadIssue(route.params.id);
    }

    return {
      isEdit,
      form,
      users,
      isCompletedOrCancelled,
      handleSubmit,
      goBack,
    };
  },
};
</script>
