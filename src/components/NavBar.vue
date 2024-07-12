<script>
// import { ref, onMounted } from 'vue'

export default {
  data() {
    return {
      useMobileNav: false,
      windowWidth: window.innerWidth
    }
  },
  methods: {
    onResize() {
      this.windowWidth = window.innerWidth
    }
  },
  mounted() {
    this.$nextTick(() => {
      if (this.windowWidth < 800){
        this.useMobileNav=true
      } else{
        this.useMobileNav = false
      }
      window.addEventListener('resize', this.onResize);
    })
  },

  beforeUnmount() { 
    window.removeEventListener('resize', this.onResize); 
  },
  // created() {
  //   window.addEventListener("resize", this.)
  // }
};

</script>

<template>
  <nav class="bg-white flex justify-between items-center fixed top-0 w-full z-10 shadow-md">
    <!-- Left-aligned link -->
    <div>
      <a href="#" class="text-gray-800 px-3 py-2 text-lg hover:bg-customBlue hover:text-white transition-colors duration-300 flex items-center">Home</a>
    </div>

    <!-- Right-aligned links -->
    <div class="flex items-center">
      <a href="#" class="text-gray-800 px-3 py-2 text-lg hover:bg-customBlue hover:text-white transition-colors duration-300 flex items-center">Link 1</a>
      <a href="#" class="text-gray-800 px-3 py-2 text-lg hover:bg-customBlue hover:text-white transition-colors duration-300 flex items-center">Link 2</a>
      <a href="#" class="text-gray-800 px-3 py-2 text-lg hover:bg-customBlue hover:text-white transition-colors duration-300 flex items-center">Link 3</a>
    </div>
  </nav>
</template>

<style scoped>
/* Custom hover effect for links */
a:hover {
  background-color: #57B8FF;
  color: #ffffff; /* Set text color to white on hover */
}

/* Ensure the link takes up the full height */
a {
  display: block;
  height: 100%;
}

/* Optionally, adjust text and link alignment */
a,
a:hover {
  text-decoration: none; /* Remove underline on hover */
  display: flex;
  align-items: center;
}
</style>