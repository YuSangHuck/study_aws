$KST = $(Get-Date -UFormat %s)
$HOUR = $(60 * 60)
$EPOCH = ($KST - (9 * $HOUR))
$INTERVAL = 2 * $HOUR

$LOG_PATH = ".\logs"
$FILE_FORMAT = $(Get-Date -Format "yyyy_MMdd_HH_mm")

If(!(test-path $LOG_PATH)) {
  New-Item -ItemType Directory -Force -Path $LOG_PATH
}


aws xray get-service-graph --start-time $(($EPOCH - $INTERVAL)) --end-time $EPOCH --profile indiv >> "$LOG_PATH\service_graph_$FILE_FORMAT.json"
aws xray get-trace-summaries --start-time $(($EPOCH - $INTERVAL)) --end-time $EPOCH --profile indiv >> "$LOG_PATH\trace_summaries_$FILE_FORMAT.json"


# Get full traces
$TRACEIDS = $(aws xray get-trace-summaries --start-time $(($EPOCH - $INTERVAL)) --end-time $(($EPOCH)) --query 'TraceSummaries[*].Id' --output text --profile indiv)
echo $TRACEIDS
aws xray batch-get-traces --trace-ids $TRACEIDS --query 'Traces[*]' --profile indiv >> "$LOG_PATH\trace_full_$FILE_FORMAT.json"

