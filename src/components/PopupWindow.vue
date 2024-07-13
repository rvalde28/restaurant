<template>
  <transition name="popup-slide">
    <div v-if="isOpen" class="flex items-center justify-center fixed top-0 right-0 bottom-0 left-0">
      <div id="windowBackground" class="bg-gray-800 bg-opacity-25 absolute top-0 right-0 bottom-0 left-0" @click="closePopupOnOutsideClick"></div> <!-- Background overlay -->

      <div class="max-w-sm sm:max-w-lg lg:max-w-2xl w-full md:w-3/4 lg:w-2/3 xl:w-1/3 bg-white shadow-lg rounded-lg transform translate-x-full xl:translate-x-0 absolute top-1/3 left-1/2 xl:custom-left-three-fifths transform -translate-x-1/2 -translate-y-1/2">
        <!-- Close button -->
        <button @click="closePopup" class="z-[1001] absolute top-0 right-0 p-2 m-2 text-gray-600 hover:text-gray-800 bg-white rounded-full shadow-md">
          <svg class="w-6 h-6 text-gray-600 hover:text-gray-800" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>

        <!-- Popup content -->
        <div class="">
          <slot></slot>
        </div>
      </div>
    </div>
  </transition>
</template>
<script>
export default {
  name: 'PopupWindow',
  props: {
    isOpen: {
      type: Boolean,
      required: true
    }
  },
  methods: {
    closePopup() {
      this.$emit('close');
    },
    closePopupOnOutsideClick(event) {
       // Check if the click target is the background overlay or element with id "windowBackground"
       if (event.target.id === 'windowBackground') {
        console.log("background")
        this.closePopup();
      }
    }
  }
};
</script>

<style scoped>
.custom-left-three-fifths{
  left: 60%;
}
.popup-slide-enter-active, .popup-slide-leave-active {
  transition: transform 0.3s ease;
}
.popup-slide-enter-from, .popup-slide-leave-to {
  transform: translateX(100%);
}
/* Centering adjustments for small screens */
@media (max-width: 770px) {
  .max-w-sm {
    width: 80%; /* Adjust the width as needed */
  }
}
</style>
