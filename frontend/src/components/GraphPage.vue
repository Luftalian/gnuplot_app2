<template>
  <div id="page">
    <div id="forms">
      <form class="form2">
        <br />
        <label for="x_range_start">X Range Start:</label>
        <input type="text" id="x_range_start" v-model="xRangeStart" />

        <label for="x_range_end">X Range End:</label>
        <input type="text" id="x_range_end" v-model="xRangeEnd" />

        <label for="y_range_start">Y Range Start:</label>
        <input type="text" id="y_range_start" v-model="yRangeStart" />

        <label for="y_range_end">Y Range End:</label>
        <input type="text" id="y_range_end" v-model="yRangeEnd" />

        <label for="plot_script">Plot Script:</label>
        <input type="text" id="plot_script" v-model="plotScript" />

        <label for="comment_request">Comment Request:</label>
        <input type="text" id="comment_request" v-model="commentRequest" />
      </form>

      <div id="image_container">
        <img :src="graphImage" alt="Graph Image" v-if="graphImage" />
        <div v-else>Loading...</div>
      </div>

      <form class="form1">
        <br />
        <label for="title">Graph Title:</label>
        <input type="text" id="title" v-model="title" />

        <label for="xLabel">X-Axis Label:</label>
        <input type="text" id="xLabel" v-model="xLabel" />

        <label for="yLabel">Y-Axis Label:</label>
        <input type="text" id="yLabel" v-model="yLabel" />

        <label for="title_size">Title Size:</label>
        <input type="number" id="title_size" v-model="titleSize" />

        <label for="image_size">Image Size:</label>
        <input type="number" id="image_size" v-model="imageSize" />

        <label for="font">Font:</label>
        <input type="text" id="font" v-model="font" />

        <label for="key_position">Key Position:</label>
        <select id="key_position" v-model="keyPosition">
          <option value="top left">top left</option>
          <option value="top right">top right</option>
          <option value="left bottom">left bottom</option>
          <option value="right bottom">right bottom</option>
        </select>
      </form>

      <form class="upload-form">
        <h2>ファイルアップロード</h2>
        <input type="file" @change="onFileChange" />
        <button type="button" @click="uploadFile">アップロード</button>
      </form>
    </div>

    <div id="tabs">
      <ul>
        <li
          v-for="(tab, index) in tabs"
          :key="index"
          @click="selectedTab = index"
          :class="{ active: selectedTab === index }"
        >
          {{ uploadedFileName }}
        </li>
      </ul>

      <file-form
        v-for="(tab, index) in tabs"
        :key="index"
        v-show="selectedTab === index"
        :index="index"
        :uploadedFileName="uploadedFileName"
        @update-file="updateFile(index, $event)"
      ></file-form>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import FileForm from "./FileForm.vue";

export default {
  components: {
    FileForm,
  },
  data() {
    return {
      uploadedFileName: "",
      title: "タイトル",
      xLabel: "x軸",
      yLabel: "y軸",
      titleSize: null,
      imageSize: null,
      font: "Arial",
      keyPosition: "top left",
      xtics: true,
      ytics: true,
      grid: false,
      xRangeStart: "",
      xRangeEnd: "",
      yRangeStart: "",
      yRangeEnd: "",
      plotScript: "",
      commentRequest: "",
      graphImage: null,
      file: null,
      tabs: [],
      selectedTab: 0,
      files: [],
    };
  },
  async mounted() {
    await this.fetchGraphImage();
  },
  watch: {
    title: "submitForm",
    xLabel: "submitForm",
    yLabel: "submitForm",
    titleSize: "submitForm",
    imageSize: "submitForm",
    font: "submitForm",
    xtics: "submitForm",
    ytics: "submitForm",
    grid: "submitForm",
    xRangeStart: "submitForm",
    xRangeEnd: "submitForm",
    yRangeStart: "submitForm",
    yRangeEnd: "submitForm",
  },
  methods: {
    updateFile(index, fileData) {
      this.$set(this.files, index, fileData);
      this.submitForm();
    },
    async submitForm() {
      const figureData = {
        title: this.title,
        title_size: parseInt(this.titleSize),
        x_label: this.xLabel,
        y_label: this.yLabel,
        image_size: parseFloat(this.imageSize),
        font: this.font,
        file: this.files,
        xtics: this.xtics,
        ytics: this.ytics,
        grid: this.grid,
        x_range_start: this.xRangeStart,
        x_range_end: this.xRangeEnd,
        y_range_start: this.yRangeStart,
        y_range_end: this.yRangeEnd,
        plot_script: this.plotScript,
        comment_request: this.commentRequest,
      };

      const graphData = {
        id: 1,
        user: "John Doe",
        user_id: 123,
        figure_data: figureData,
        script: "print('Script executed successfully.')",
        comment_request: this.commentRequest,
      };
      await this.sendGraphData(graphData);
    },
    async fetchGraphImage() {
      try {
        const baseUrl = process.env.VUE_APP_BACKEND_URL;
        const response = await fetch(`${baseUrl}/api/v2/files/edit`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({}),
        });
        if (response.ok) {
          const blob = await response.blob();
          this.graphImage = URL.createObjectURL(blob);
        } else {
          console.error("Failed to fetch graph image");
        }
      } catch (error) {
        console.error("Error fetching graph image:", error);
      }
    },
    async sendGraphData(data) {
      try {
        this.graphImage = null;
        const baseUrl = process.env.VUE_APP_BACKEND_URL;
        const response = await fetch(`${baseUrl}/api/v2/files/edit`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        });
        if (response.ok) {
          const blob = await response.blob();
          this.graphImage = URL.createObjectURL(blob);
        } else {
          console.error("Failed to fetch graph image");
        }
      } catch (error) {
        console.error("Error fetching graph image:", error);
      }
    },
    onFileChange(event) {
      this.file = event.target.files[0];
    },
    async uploadFile() {
      if (!this.file) {
        alert("ファイルを選択してください。");
        return;
      }

      const formData = new FormData();
      formData.append("file", this.file);

      try {
        const baseUrl = process.env.VUE_APP_BACKEND_URL;
        const response = await axios.post(`${baseUrl}/api/v2/upload`, formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });
        alert(response.data.message);
        if (response.status === 200) {
          const uploadedFile = formData.get("file");
          if (uploadedFile) {
            const fileData = {
              data_file_name: uploadedFile.name,
              key_value: "Data",
              points_or_line: "points",
              points_size: 5,
              point_type: "1",
              line_width: "2",
              plot_line_color: "black",
            };
            this.files.push(fileData);
            this.tabs.push({});
            this.selectedTab = this.tabs.length - 1;
            this.uploadedFileName = uploadedFile.name;
          }
        }
      } catch (error) {
        console.error("ファイルアップロード中にエラーが発生しました:", error);
        alert("ファイルアップロード中にエラーが発生しました。");
      }
    },
  },
};
</script>

<style>
/* フォーム全体のスタイル */
form {
  display: grid;
  gap: 10px;
  max-width: 600px;
  margin: auto;
}

/* ラベルのスタイル */
label {
  font-weight: bold;
}

/* 入力要素のスタイル */
input[type="text"],
input[type="number"],
select {
  width: 100%;
  padding: 8px;
  border-radius: 5px;
  border: 1px solid #ccc;
}

/* 画像のスタイル */
img {
  max-width: 100%;
  height: auto;
  margin-top: 10px;
}
div {
  text-align: center;
}

/* フォーム内の要素の配置 */
form > * {
  grid-column: 1 / -1;
}

/* 親要素のスタイル */
div {
  display: flex;
  justify-content: center; /* 要素の間隔を均等に */
  align-items: flex-start; /* 上揃え */
  width: 100%;
}

form {
  margin-top: 1%;
}

/* form1のスタイル */
.form1 {
  width: calc(50% - 5px); /* 幅を半分に */
  max-width: 400px; /* 最大幅 */
  margin: 5%;
  margin-top: 0;
}

/* グラフ画像のスタイル */
img {
  width: 100%;
}

#image_container {
  width: 600px; /* 固定幅 */
  height: auto;
  border: 1px solid #ccc; /* 境界線を追加 */
  padding: 10px;
  box-sizing: border-box; /* パディングを幅に含める */
  display: flex;
  justify-content: center; /* コンテンツを中央揃え */
  align-items: center; /* コンテンツを中央揃え */
}

#image_container div {
  width: 100%; /* ローディングメッセージの幅を親に合わせる */
  text-align: center; /* ローディングメッセージを中央揃え */
}

/* タブのスタイル */
#tabs {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  background-color: #ccc;
}

ul {
  list-style: none;
  padding: 0;
  display: flex;
  justify-content: flex-start;
  margin: 0;
  /* border-bottom: 1px solid #ccc; */
  background-color: #ccc;
}

li {
  margin: 0;
  padding: 10px 20px;
  cursor: pointer;
  background-color: #ccc; /* 未選択のタブの背景色 */
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
  border: 1px solid #bbb;
  border-bottom: none;
}

li.active {
  background-color: #fff; /* 選択されているタブの背景色 */
}

/* 値の表示エリアのスタイル */
.file-form-container {
  flex: 1;
  padding: 20px;
  border: 1px solid #ccc;
  margin-top: 10px;
  background-color: #fff;
}

#page {
  display: flex;
  flex-direction: column;
}
</style>
