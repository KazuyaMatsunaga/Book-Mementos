<template>
  <div class="list">
    <div class="container">
      <img class="list-logo" src="http://localhost:8080/assets/img/list-gopher.png/" width="325px" />
      <div class="book-wrapper">
        <b-card-group deck="width: 100%">
          <div class="book-card" v-for="book in books" :key="book">
            <b-card no-body class="overflow-hidden" border-variant="dark" style="width: 540px">
              <b-row no-gutters>
                <b-col md="6">
                  <b-card-img v-bind:src="book.ImageLink"></b-card-img>
                </b-col>
                <b-col md="6">
                  <b-card-body v-bind:title="book.Title" v-bind:sub-title="book.SubTitle">
                    <b-card-text>
                      {{ book.Authors }}
                      <br />
                      {{ book.Publisher }}
                      <br />
                      {{ book.PublishedDate }}
                    </b-card-text>
                    <b-button
                      variant="outline-dark"
                      v-bind:href="book.PreviewLink"
                      target="_blank"
                    >Show Detail</b-button>
                  </b-card-body>
                </b-col>
              </b-row>
            </b-card>
          </div>
        </b-card-group>
      </div>
    </div>
  </div>
</template>

<script>
module.exports = {
  data() {
    return {
      // 本の情報
      books: []
    };
  },
  // インスタンス作成時の処理
  created: function() {
    this.doFetchAllBooksFromList();
  },
  methods: {
    doFetchAllBooksFromList() {
      axios.get("/fetchAllBooksFromList").then(response => {
        if (response.status != 200) {
          console.log(response.status);
        } else {
          for (let i = 0; i < response.data.length; i++) {
            this.books.push(response.data[i]);
          }
          console.log(this.books);
        }
      });
    }
  }
};
</script>

<style>
.list-logo {
  display: block;
  margin: 0 auto;
}

.book-wrapper {
  width: 100%;
  margin-top: 30px;
  margin-bottom: 30px;
  display: block;
  margin: 0 auto;
}

.book-card {
  margin-top: 30px;
  position: relative;
}
</style>