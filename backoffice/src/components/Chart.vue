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
      this.render()
    }
  },
  methods: {
    render () {
      this.renderChart({
        labels: this.chartLabels,
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
    'chartData',
    'chartLabels'
  ])
})
</script>
