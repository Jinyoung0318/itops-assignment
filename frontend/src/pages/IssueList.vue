<template>
  <div class="issue-list">
    <h1>이슈 목록</h1>
    <div>
      <label for="status">상태 필터:</label>
      <select id="status" v-model="selectedStatus" @change="fetchIssues">
        <option value="">전체</option>
        <option value="PENDING">PENDING</option>
        <option value="IN_PROGRESS">IN_PROGRESS</option>
        <option value="COMPLETED">COMPLETED</option>
        <option value="CANCELLED">CANCELLED</option>
      </select>
      <button @click="navigateToCreate">새 이슈 생성</button>
    </div>
    <ul>
      <li v-for="issue in issues" :key="issue.id" @click="goToDetail(issue.id)">
        <h3>{{ issue.title }}</h3>
        <p>상태: {{ issue.status }}</p>
        <p>담당자: {{ issue.user?.name || '없음' }}</p>
        <p>생성일: {{ formatDate(issue.createdAt) }}</p>
      </li>
    </ul>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'IssueList',
  setup() {
    const issues = ref([])
    const selectedStatus = ref('')
    const router = useRouter()

    const fetchIssues = async () => {
      const query = selectedStatus.value ? `?status=${selectedStatus.value}` : ''
      const res = await fetch(`http://localhost:8080/issues${query}`)
      const data = await res.json()
      issues.value = data.issues
    }

    const goToDetail = (id) => {
      router.push(`/issue/${id}`)
    }

    const navigateToCreate = () => {
      router.push('/issue/new')
    }

    const formatDate = (dateStr) => {
      const date = new Date(dateStr)
      return date.toLocaleString()
    }

    onMounted(fetchIssues)

    return {
      issues,
      selectedStatus,
      fetchIssues,
      goToDetail,
      navigateToCreate,
      formatDate
    }
  }
}
</script>
