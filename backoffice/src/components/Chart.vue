<script>
import { Line } from 'vue-chartjs'
import { mapGetters } from 'vuex'

export default Line.extend({
  props: ['labels'],
  mounted () {
    this.render()
  },
  watch: {
    datasetLabel (value) {
      this.render()
    },
    chartData (value) {
      console.log(value)
      this.render()
    }
  },
  methods: {
    render () {
      this.renderChart({
        labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
        datasets: [
          {
            label: this.datasetLabel,
            backgroundColor: '#FFF',
            data: this.chartData
          }
        ],
        responsive: true
      })
    }
  },
  computed: mapGetters([
    'datasetLabel',
    'chartData'
  ])
})
</script>
