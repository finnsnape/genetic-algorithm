<script lang="ts">
  import type { Image } from "$lib/ts/types";
  import Slide from "$lib/components/Slide.svelte";

  const slides = [
    {
      title: "Crossover",
      description: [
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.", 
        "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
      ],
      image: {
        alt: "", 
        src: "https://media.geeksforgeeks.org/wp-content/uploads/20190620121215/singleCrossover-2.png"
      }
    },
    {
      title: "Mutation",
      description: [
        "Aorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.", 
        "Vuis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
      ],
      image: {
        alt: "", 
        src: "https://media.geeksforgeeks.org/wp-content/uploads/20190620121215/singleCrossover-2.png"
      }
    },
    {
      title: "Breeding",
      description: [
        "Oorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.", 
        "Luis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
      ],
      image: {
        alt: "", 
        src: "https://media.geeksforgeeks.org/wp-content/uploads/20190620121215/singleCrossover-2.png"
      }
    }, 
  ]

  let slideIndex: number = 0;

  function incrementSlideIndex() {
    slideIndex;
  }
</script>

<Slide title={slides[slideIndex].title} description={slides[slideIndex].description} image={slides[slideIndex].image}></Slide>

<button on:click={incrementSlideIndex}>Next</button>