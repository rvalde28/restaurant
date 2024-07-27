<template>
    <div class="mt-12 z-0">
      <h1>About Page</h1>
      <button @click="navigateToHome">Go to Home Page</button>

    </div>
</template>

<script>
import NavBar from '../components/NavBar.vue';

export default {
    components: {
        NavBar
    },
    methods: {
        navigateToHome() {
        this.$router.push('/');
        }
    }
};
</script>

<style>
.cart-bar{
  z-index: 100;
}
/* .menu-app{
  overflow:scroll;
  height:auto;
  -webkit-overflow-scrolling: touch;
} */
.shopping-button {
  position: fixed;
  bottom: 12px;
  right: 0;
  border-radius: 50%;
  display: flex;
  width: 60px;
  height: 60px;
}
</style>
