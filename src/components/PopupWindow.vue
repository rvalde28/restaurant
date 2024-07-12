<template>

  <transition name="popup-slide">
    <div v-if="isOpen" class="fixed top-0 right-0 bottom-0 left-0 flex items-center justify-center">
      <div id="windowBackground" v-if="isOpen" class="bg-gray-800 bg-opacity-50 fixed top-0 right-0 bottom-0 left-0" @click="closePopupOnOutsideClick"></div> <!-- Background overlay -->

      <div class="right-space max-w-sm w-full md:w-3/4 lg:w-2/3 xl:w-1/3 bg-white shadow-lg rounded-lg p-4 transform translate-x-full md:translate-x-0 absolute top-1/2 transform -translate-y-1/2">
        <!-- Close button -->
        <button @click="closePopup" class="absolute top-0 right-0 p-2 m-4 text-gray-600 hover:text-gray-800">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>

        <!-- Popup content -->
        <div class="py-2">
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
.right-space{
  z-index: 1000000000000000; /* Adjust the value as needed */
  right: 10vw;
  top: 30vh;
}

.popup-slide-enter-active, .popup-slide-leave-active {
  transition: transform 0.3s ease;
}
.popup-slide-enter-from, .popup-slide-leave-to {
  transform: translateX(100%);
}
</style>
