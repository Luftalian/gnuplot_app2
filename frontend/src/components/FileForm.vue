<template>
  <div id="data_files">
    <div>
      <label for="data_file_name">Data File Name:</label>
      <input type="text" id="data_file_name" v-model="dataFileName" />

      <label for="key_value">Key Value</label>
      <input type="text" id="key_value" v-model="keyValue" />

      <label for="plot_line_color">Plot Line Color:</label>
      <select id="plot_line_color" v-model="plotLineColor">
        <option value="black">black</option>
        <option value="red">red</option>
        <option value="blue">blue</option>
        <option value="green">green</option>
      </select>
    </div>

    <div>
      <label for="points_or_line">Points or Line:</label>
      <select id="points_or_line" v-model="pointsOrLine">
        <option value="points">points</option>
        <option value="lines">lines</option>
        <option value="linespoints">linespoints</option>
      </select>

      <label for="points_size">Points Size:</label>
      <input type="number" id="points_size" v-model="pointsSize" />

      <label for="point_type">Point Type:</label>
      <select id="point_type" v-model="pointType">
        <option value="1">circle</option>
        <option value="2">square</option>
        <option value="3">triangle</option>
        <option value="4">cross</option>
        <option value="5">plus</option>
        <option value="6">star</option>
        <option value="7">x</option>
      </select>

      <label for="line_width">Line Width:</label>
      <input type="text" id="line_width" v-model="lineWidth" />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    index: Number, // indexの型を明示的に指定
    uploadedFileName: String, // 追加: アップロードされたファイル名を受け取る
  },
  data() {
    return {
      dataFileName: this.uploadedFileName, // 追加: アップロードされたファイル名を受け取る
      keyValue: "",
      pointsOrLine: "points",
      pointsSize: 1,
      pointType: "1",
      lineWidth: "",
      plotLineColor: "black",
    };
  },
};
</script>

<style>
div {
  background-color: #fff;
}
.file-form {
  display: grid;
  gap: 10px;
  grid-template-columns: 1fr 2fr; /* 1つ目の列はラベル、2つ目の列は入力フィールド */
  margin-bottom: 20px;
}

label {
  font-weight: bold;
}

#data_files {
  padding-top: 3%;
  display: flex;
  flex-direction: column;
  text-align: center;
  justify-content: center; /* 要素の間隔を均等に */
  align-items: flex-start; /* 上揃え */
}

input[type="text"],
input[type="number"],
select {
  width: 100%;
  padding: 8px;
  border-radius: 5px;
  border: 1px solid #ccc;
}

select {
  appearance: none; /* デフォルトのスタイルを削除 */
}

select::-ms-expand {
  display: none; /* IE11の矢印アイコンを非表示 */
}
</style>
