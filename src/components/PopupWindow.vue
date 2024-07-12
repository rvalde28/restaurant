<template>
    <transition name="popup-slide">
      <div v-if="isOpen" class="fixed top-0 right-0 bottom-0 left-0 flex items-center justify-center bg-gray-800 bg-opacity-50">
        <div class="max-w-lg w-full bg-white shadow-lg rounded-lg p-6">
          <!-- Close button -->
          <button @click="closePopup" class="absolute top-0 right-0 p-2 m-4 text-gray-600 hover:text-gray-800">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
  
          <!-- Popup content -->
          <div>
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
      }
    }
  };
  </script>
  
  <style scoped>
  .popup-slide-enter-active, .popup-slide-leave-active {
    transition: transform 0.3s ease;
  }
  .popup-slide-enter-from, .popup-slide-leave-to {
    transform: translateX(100%);
  }
  </style>