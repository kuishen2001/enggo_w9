<!-- App.svelte -->
<script>
  import logo from './assets/images/GoQuery.png'
  import { Query, Search } from '../wailsjs/go/main/App.js'

  let resultText = "Please enter your query below 👇"
  let queryterm = "";
  let queryInfo = "";

  async function search() {
    console.log("Search button clicked");
    try {
      const result = await Query(queryterm);
      resultText = result;
    } catch (error) {
      resultText = "An error occurred: " + error.message;
      console.error("Error in search:", error);
    }
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src={logo}>
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={queryterm} class="input" id="queryterm" type="text"/>
  </div>
  <button class="btn" on:click={search}>Search</button>
  <div class="query-info" id="query-info">
    <pre>{queryInfo}</pre>
  </div>
</main>

<style>

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
