import {
  ObservabilityVendor,
  ObservabilitySignals,
  VendorObjects,
  VendorType,
} from "@/vendors/index";
import CoralogixLogo from "@/img/vendor/coralogix.svg";
import { NextApiRequest } from "next";

export class Coralogix implements ObservabilityVendor {
  name = "coralogix";
  displayName = "Coralogix";
  type = VendorType.MANAGED;
  supportedSignals = [
    ObservabilitySignals.Traces,
    ObservabilitySignals.Metrics,
    ObservabilitySignals.Logs,
  ];

  getLogo = (props: any) => {
    return <CoralogixLogo {...props} />;
  };

  getFields = (selectedSignals: any) => {
    return [
      {
        displayName: "Logs Endpoint",
        id: "logsEndpoint",
        name: "logsEndpoint",
        type: "url",
      },
      {
        displayName: "Metrics Endpoint",
        id: "metricsEndpoint",
        name: "metricsEndpoint",
        type: "url",
      },
      {
        displayName: "Traces Endpoint",
        id: "tracesEndpoint",
        name: "tracesEndpoint",
        type: "url",
      },
      {
        displayName: "Application Name",
        id: "appName",
        name: "appName",
        type: "text",
      },
      {
        displayName: "Subsystem Name",
        id: "subsystemName",
        name: "subsystemName",
        type: "text",
      },
      {
        displayName: "API Key",
        id: "apikey",
        name: "apikey",
        type: "password",
      },
    ];
  };

  toObjects = (req: NextApiRequest) => {
    return {
      Data: {
        CORALOGIX_LOGS_ENDPOINT: req.body.logsEndpoint,
        CORALOGIX_METRICS_ENDPOINT: req.body.metricsEndpoint,
        CORALOGIX_TRACES_ENDPOINT: req.body.tracesEndpoint,
        CORALOGIX_APPNAME: req.body.appName,
        CORALOGIX_SUBSYSTEMNAME: req.body.subsystemName,
      },
      Secret: {
        CORALOGIX_API_KEY: Buffer.from(req.body.apikey).toString("base64"),
      },
    };
  };

  mapDataToFields = (data: any) => {
    return {
      logsEndpoint: data.CORALOGIX_LOGS_ENDPOINT,
      metricsEndpoint: data.CORALOGIX_METRICS_ENDPOINT,
      tracesEndpoint: data.CORALOGIX_TRACES_ENDPOINT,
      appName: data.CORALOGIX_APPNAME,
      subsystemName: data.CORALOGIX_SUBSYSTEMNAME,
    };
  };
}
