package gosqlprometheus

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type sqlCollector struct {
	db *sql.DB

	maxOpenConnections                                 *prometheus.Desc
	openConnections, inUseConnections, idleConnections *prometheus.Desc
	waitCount, waitDuration                            *prometheus.Desc
	maxIdleClosed, maxLifetimeClosed                   *prometheus.Desc
}

// SQLCollectorOpts defines the behavior of a sql metrics collector
// created with NewSQLCollector.
type SQLCollectorOpts struct {
	DBName string
	DB     *sql.DB
}

func NewSQLCollector(opts SQLCollectorOpts) prometheus.Collector {
	if opts.DB == nil {
		panic("DB can not be nil")
	}
	labels := make(map[string]string)
	if len(opts.DBName) > 0 {
		labels["db"] = opts.DBName
	}

	c := &sqlCollector{
		db: opts.DB,
		maxOpenConnections: prometheus.NewDesc(
			"sql_max_open_connections",
			"Maximum number of open connections to the database.",
			nil, labels,
		),
		openConnections: prometheus.NewDesc(
			"sql_open_connections",
			"The number of established connections both in use and idle.",
			nil, labels,
		),
		inUseConnections: prometheus.NewDesc(
			"sql_in_use_connections",
			"The number of connections currently in use.",
			nil, labels,
		),
		idleConnections: prometheus.NewDesc(
			"sql_idle_connections",
			"The number of idle connections.",
			nil, labels,
		),
		waitCount: prometheus.NewDesc(
			"sql_wait_count",
			"The total number of connections waited for.",
			nil, labels,
		),
		waitDuration: prometheus.NewDesc(
			"sql_wait_duration",
			"The total time blocked waiting for a new connection.",
			nil, labels,
		),
		maxIdleClosed: prometheus.NewDesc(
			"sql_max_idle_closed",
			"The total number of connections closed due to SetMaxIdleConns.",
			nil, labels,
		),
		maxLifetimeClosed: prometheus.NewDesc(
			"sql_max_lifetime_closed",
			"The total number of connections closed due to SetMaxIdleConns.",
			nil, labels,
		),
	}

	return c
}

// Describe returns all descriptions of the collector.
func (c *sqlCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.maxOpenConnections
	ch <- c.openConnections
	ch <- c.inUseConnections
	ch <- c.idleConnections
	ch <- c.waitCount
	ch <- c.waitDuration
	ch <- c.maxIdleClosed
	ch <- c.maxLifetimeClosed
}

// Collect returns the current state of all metrics of the collector.
func (c *sqlCollector) Collect(ch chan<- prometheus.Metric) {
	stats := c.db.Stats()
	ch <- prometheus.MustNewConstMetric(c.maxOpenConnections, prometheus.GaugeValue, float64(stats.MaxOpenConnections))
	ch <- prometheus.MustNewConstMetric(c.openConnections, prometheus.GaugeValue, float64(stats.OpenConnections))
	ch <- prometheus.MustNewConstMetric(c.inUseConnections, prometheus.GaugeValue, float64(stats.InUse))
	ch <- prometheus.MustNewConstMetric(c.idleConnections, prometheus.GaugeValue, float64(stats.Idle))
	ch <- prometheus.MustNewConstMetric(c.waitCount, prometheus.CounterValue, float64(stats.WaitCount))
	ch <- prometheus.MustNewConstMetric(c.waitDuration, prometheus.CounterValue, float64(stats.WaitDuration))
	ch <- prometheus.MustNewConstMetric(c.maxIdleClosed, prometheus.CounterValue, float64(stats.MaxIdleClosed))
	ch <- prometheus.MustNewConstMetric(c.maxLifetimeClosed, prometheus.CounterValue, float64(stats.MaxLifetimeClosed))
}
