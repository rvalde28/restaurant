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
    <div class="top-nav" v-if="!useMobileNav" >
      <div class="left-links">
         <a class="links active home center-text-vertically" href="#">Home</a>
      </div>
      <div class="right-links">
        <a class="links center-text-vertically" href="#"> Order Online</a>
        <a class="links center-text-vertically" href="#"> Menu</a>
        <a class="links center-text-vertically" href="#"> Contact Us</a>
        <a class="links center-text-vertically" href="#">Login</a>
      </div>
    </div>

    <div class="mobile-nav" v-if="useMobileNav" >
      <a href="#home">Home</a>
      <a href="#news">News</a>
      <a href="#contact">Contact</a>
      <a href="#about">About</a>
      <a href="javascript:void(0);" class="icon" onclick="myFunction()">&#9776;</a>
    </div>
</template>

<style scoped>
a {
    height: 100%;
}

.mobile-nav {
  z-index: 10000;
  background-color: #333;
  overflow: none;
  position: fixed;
  bottom: 0;
  width: 100%;
}

/* Style the links inside the navigation bar */
.mobile-nav a {
  float: left;
  display: block;
  color: #f2f2f2;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}

.top-nav{
  z-index: 10000;
  background-color: white;
  height: 100%;
  display: flex;
  position: fixed;
  top:0;
  width: 100%;
  overflow: none;
  height: auto;
  color: gray;
}

.top-nav a.active {
    border-bottom: 3px solid #57B8FF;
}

.left-links{
    flex:1 1 200px;
}

.links {
    font-weight: 600;
    float: left;
    color: gray;
    text-align: center;
    height: 100%;
    padding-left: 14px;
    padding-right: 14px;
    text-decoration: none;
    font-size: 17px;
}

a:hover {
    background: #57B8FF;
    color: white;
}
a:selected{
    background:#cdedfd;
}



@media (min-width: 360px) {
  home{

  }
  .links {
    font-size: .9rem;
  }
}

@media (min-width: 512px) {
  .links {
    font-size: .9rem;
  }
}

@media (min-width: 720px) {
  .links {
    font-size: .9rem;
  }
}
@media (min-width: 1080px) {
  .links {
    font-size: 1.2vw;
  }
  .left-links{
    height: 2.9vw;  /* has to be the same as height and line-height in center-text-vertically */
  }
  .center-text-vertically{
    height: 2.9vw;       
    line-height: 2.9vw;  /* has to be the same as height */
    text-align: center;
    /* border: 2px dashed #f69c55; */
  }
  
}

</style>