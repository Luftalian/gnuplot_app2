<template>
  <div>
    <form class="form2">
      <br>

      <br>
      <label for="x_range_start">X Range Start:</label>
      <input type="text" id="x_range_start" v-model="x_range_start" />

      <label for="x_range_end">X Range End:</label>
      <input type="text" id="x_range_end" v-model="x_range_end" />

      <label for="y_range_start">Y Range Start:</label>
      <input type="text" id="y_range_start" v-model="y_range_start" />

      <label for="y_range_end">Y Range End:</label>
      <input type="text" id="y_range_end" v-model="y_range_end" />

      <label for="plot_script">Plot Script:</label>
      <input type="text" id="plot_script" v-model="plot_script" />

      <label for="comment_request">Comment Request:</label>
      <input type="text" id="comment_request" v-model="comment_request" />

      <label for="data_file_name">Data File Name:</label>
      <input type="text" id="data_file_name" v-model="data_file_name" />
    </form>
    <div id="image_container">
      <img :src="graphImage" alt="Graph Image" v-if="graphImage">
      <div v-else>Loading...</div>
    </div>
    <form class="form1">
      <br>
      <label for="title">Graph Title:</label>
      <input type="text" id="title" v-model="title" />

      <label for="xLabel">X-Axis Label:</label>
      <input type="text" id="xLabel" v-model="xLabel" />

      <label for="yLabel">Y-Axis Label:</label>
      <input type="text" id="yLabel" v-model="yLabel" />

      <label for="title_size">title_size:</label>
      <input type="number" id="title_size" v-model="title_size" />

      <label for="image_size">image_size:</label>
      <input type="number" id="image_size" v-model="image_size" />

      <!-- default is 'Arial' -->
      <label for="font">Font:</label>
      <input type="text" id="font" v-model="font" />
      <label for="key_value">Key Value</label>
      <input type="text" id="key_value" v-model="key_value" />
      <label for="key_position">Key Position:</label>
      <select id="key_position" v-model="key_position">
        <option value="top left">top left</option>
        <option value="top right">top right</option>
        <option value="left bottom">left bottom</option>
        <option value="right bottom">right bottom</option>
      </select>

      <label for="points_or_line">Points or Line:</label>
      <select id="points_or_line" v-model="points_or_line">
        <option value="points">points</option>
        <option value="lines">lines</option>
        <option value="linespoints">linespoints</option>
      </select>

      <label for="points_size">Points Size:</label>
      <input type="number" id="points_size" v-model="points_size" />

      <label for="point_type">Point Type:</label>
      <select id="point_type" v-model="point_type">
        <option value="1">circle</option>
        <option value="2">square</option>
        <option value="3">triangle</option>
        <option value="4">cross</option>
        <option value="5">plus</option>
        <option value="6">star</option>
        <option value="7">x</option>
      </select>

      <label for="line_width">Line Width:</label>
      <input type="text" id="line_width" v-model="line_width" />

      <label for="plot_line_color">Plot Line Color:</label>
      <select id="plot_line_color" v-model="plot_line_color">
        <option value="black">black</option>
        <option value="red">red</option>
        <option value="blue">blue</option>
        <option value="green">green</option>
      </select>

      <br>

      <label for="xtics">X Ticks:</label>
      <input type="checkbox" id="xtics" v-model="xtics" />

      <br>

      <label for="ytics">Y Ticks:</label>
      <input type="checkbox" id="ytics" v-model="ytics" />

      <br>

      <label for="grid">Grid:</label>
      <input type="checkbox" id="grid" v-model="grid" />
    </form>

    <form class="upload-form">
      <h2>ファイルアップロード</h2>
      <input type="file" @change="onFileChange" />
      <button type="button" @click="uploadFile">アップロード</button>
    </form>

  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      title: 'タイトル',
      data_file_name: '',
      xLabel: 'x軸',
      yLabel: 'y軸',
      title_size: '',
      image_size: '',
      font: 'Arial',
      key_value: 'Data',
      key_position: 'top right',
      points_or_line: 'points',
      points_size: '',
      point_type: '1',
      line_width: '',
      plot_line_color: 'black',
      xtics: true,
      ytics: true,
      grid: false,
      x_range_start: '',
      x_range_end: '',
      y_range_start: '',
      y_range_end: '',
      plot_script: '',
      comment_request: '',
      graphImage: null,
      file: null
    };
  },
  async mounted() {
    await this.fetchGraphImage();
  },
  watch: {
    // 監視対象のデータ
    title: 'submitForm',
    data_file_name: 'submitForm',
    xLabel: 'submitForm',
    yLabel: 'submitForm',
    title_size: 'submitForm',
    image_size: 'submitForm',
    font: 'submitForm',
    key_value: 'submitForm',
    key_position: 'submitForm',
    points_or_line: 'submitForm',
    points_size: 'submitForm',
    point_type: 'submitForm',
    line_width: 'submitForm',
    plot_line_color: 'submitForm',
    xtics: 'submitForm',
    ytics: 'submitForm',
    grid: 'submitForm',
    x_range_start: 'submitForm',
    x_range_end: 'submitForm',
    y_range_start: 'submitForm',
    y_range_end: 'submitForm',
    plot_script: 'submitForm',
    comment_request: 'submitForm'
  },
  methods: {
    async submitForm() {
      const graphData = {
        id: 1,
        user: 'John Doe',
        user_id: 123,
        figure_data: {
          title: this.title,
          data_file_name: this.data_file_name,
          title_size: parseInt(this.title_size),
          x_label: this.xLabel,
          y_label: this.yLabel,
          image_size: parseFloat(this.image_size),
          font: this.font,
          key_value: this.key_value,
          key_position: this.key_position,
          points_or_line: this.points_or_line,
          points_size: parseInt(this.points_size),
          point_type: this.point_type,
          line_width: this.line_width,
          plot_line_color: this.plot_line_color,
          xtics: this.xtics,
          ytics: this.ytics,
          grid: this.grid,
          x_range_start: this.x_range_start,
          x_range_end: this.x_range_end,
          y_range_start: this.y_range_start,
          y_range_end: this.y_range_end,
          plot_script: this.plot_script
        },
        script: "print('Script executed successfully.')",
        comment_request: this.comment_request
      };
      await this.sendGraphData(graphData);
    },
    async fetchGraphImage() {
      try {
        const baseUrl = process.env.VUE_APP_BACKEND_URL;
        const response = await fetch(`${baseUrl}/api/v2/files/edit`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({})
        });
        if (response.ok) {
          const blob = await response.blob();
          this.graphImage = URL.createObjectURL(blob);
        } else {
          console.error('Failed to fetch graph image');
        }
      } catch (error) {
        console.error('Error fetching graph image:', error);
      }
    },
    async sendGraphData(data) {
      try {
        this.graphImage = null;
        const baseUrl = process.env.VUE_APP_BACKEND_URL;
        const response = await fetch(`${baseUrl}/api/v2/files/edit`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
        });
        if (response.ok) {
          const blob = await response.blob();
          this.graphImage = URL.createObjectURL(blob);
        } else {
          console.error('Failed to fetch graph image');
        }
      } catch (error) {
        console.error('Error fetching graph image:', error);
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
            'Content-Type': 'multipart/form-data',
          },
        });
        alert(response.data.message);
        if (response.status === 200) { // response.ok は Fetch API 用のプロパティ。axios では response.status で確認します。
          // formDataからファイル名を取得
          const uploadedFile = formData.get('file'); // 'file' は formData に追加したときのフィールド名に置き換えてください
          if (uploadedFile) {
            this.data_file_name = uploadedFile.name;
          }
        }
      } catch (error) {
        console.error("ファイルアップロード中にエラーが発生しました:", error);
        alert("ファイルアップロード中にエラーが発生しました。");
      }
    }
  }
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

/* ローディングメッセージのスタイル */
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
}

form {
  margin-top: 1%;
}

/* form1のスタイル */
.form1 {
  width: calc(10% - 5px); /* 幅を半分に */
  max-width: 400px; /* 最大幅 */
  margin: 5%;
  margin-top: 0;
}

/* form2のスタイル */
.form2 {
  width: calc(10% - 5px); /* 幅を半分に */
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
</style>