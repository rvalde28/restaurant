<template>
    <div>
      <h1>About Page</h1>
      <button @click="navigateToHome">Go to Home Page</button>
    </div>
</template>

<script>
    export default {
    methods: {
        navigateToHome() {
        this.$router.push('/');
        }
    }
    };
</script>