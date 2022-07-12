<script lang="ts">
  import { Router, Route } from "svelte-navigator";
  
  import Header from "./component/Header.svelte";
  import Footer from "./component/Footer.svelte";
  import Home from "./component/Home.svelte";
  import CustomerEdit from "./client/ClientEdit.svelte";
  import CustomerList from "./client/ClientList.svelte";
</script>

<style>
main {

}

</style>
<main>
  <Router>
    <Header></Header>
    <Route path="/">
      <Home/>
    </Route> 
    <Route path="client/*">
      <Route path="/">
        <CustomerList></CustomerList>
      </Route>
      <Route path="/:id" component={CustomerEdit} />
    </Route>
    <Footer></Footer>
  </Router>
</main>
