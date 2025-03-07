import type { BaseWorkflowProps, NoteConfig, WorkflowConfiguration } from "@clutch-sh/core";
import type { WizardChild } from "@clutch-sh/wizard";

import CordonNode from "./cordon-node";
import DeletePod from "./delete-pod";
import KubeDashboard from "./k8s-dashboard";
import ResizeHPA from "./resize-hpa";

interface ResolverConfigProps {
  resolverType: string;
}

interface ConfirmConfigProps {
  notes?: NoteConfig[];
}

export interface WorkflowProps extends BaseWorkflowProps, ResolverConfigProps, ConfirmConfigProps {}
export interface ResolverChild extends WizardChild, ResolverConfigProps {}
export interface ConfirmChild extends WizardChild, ConfirmConfigProps {}

const register = (): WorkflowConfiguration => {
  return {
    developer: {
      name: "Lyft",
      contactUrl: "mailto:hello@clutch.sh",
    },
    path: "k8s",
    group: "K8s",
    displayName: "K8s",
    routes: {
      deletePod: {
        path: "pod/delete",
        displayName: "Delete Pod",
        description: "Delete a K8s pod.",
        component: DeletePod,
        requiredConfigProps: ["resolverType"],
      },
      resizeHPA: {
        path: "hpa/resize",
        displayName: "Resize HPA",
        description: "Resize a horizontal autoscaler.",
        component: ResizeHPA,
        requiredConfigProps: ["resolverType"],
      },
      kubeDashboard: {
        path: "dashboard",
        displayName: "Kubernetes Dashboard",
        description: "Dashboard for Kubernetes Resources.",
        component: KubeDashboard,
        requiredConfigProps: [],
      },
      cordonNode: {
        path: "node/cordon",
        displayName: "Cordon/Uncordon Node",
        description: "Cordon or uncordon a node",
        component: CordonNode,
        requiredConfigProps: ["resolverType"],
      },
    },
  };
};

export default register;
