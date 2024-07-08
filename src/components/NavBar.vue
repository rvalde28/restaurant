<script setup>
</script>

<template>
    <div class="top-nav">
      <div class="left-links">
         <a class="links active" href="#">Home</a>
      </div>
      <div class="right-links">
        <a class="links" href="#"> Order Online</a>
        <a class="links" href="#"> Menu</a>
        <a class="links" href="#"> Contact Us</a>
        <a class="links" href="#">Login</a>
      </div>
    </div>
</template>

<style scoped>
a {
    height: 100%;
}
.top-nav{
    background-color: white;
    height: 100%;
    display: flex;
    position: fixed;
    top:0;
    width: 100%;
    overflow: auto;
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
    padding-left: 14px;
    padding-right: 14px;
    padding-top: 8px;
    padding-bottom: 8px;
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


@media (min-width: 200px) {
  .links {
    font-size: .9rem;
  }
}


@media (min-width: 600px) {
  .links {
    font-size: .9rem;
  }
}

@media (min-width: 1024px) {
  .links {
    font-size: .9rem;
  }
}

</style>