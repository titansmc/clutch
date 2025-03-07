import type { Route, Workflow } from "../AppProvider/workflow";

interface GroupedRoutes {
  [category: string]: {
    workflows: {
      displayName: string;
      path: string;
      trending: boolean;
    }[];
  };
}

interface TrendingWorkflow {
  displayName: string;
  group: string;
  description: string;
  path: string;
}

const getDisplayName = (workflow: Workflow, route: Route, delimiter: string = ":"): string => {
  let { displayName } = workflow;
  if (route.displayName) {
    displayName = `${
      displayName.toLowerCase() !== workflow.group.toLowerCase()
        ? `${displayName}${delimiter} `
        : ""
    }${route.displayName}`;
  }

  return displayName;
};

const workflowsByTrending = (workflows: Workflow[]): TrendingWorkflow[] => {
  const trending = [];

  workflows.forEach(workflow => {
    workflow.routes.forEach(route => {
      if (route.trending) {
        trending.push({
          displayName: getDisplayName(workflow, route),
          group: workflow.group,
          description: route.description,
          path: `${workflow.path}/${route.path}`,
        });
      }
    });
  });

  return trending;
};

const routesByGrouping = (workflows: Workflow[]): GroupedRoutes => {
  const routes = {};
  workflows.forEach(workflow => {
    const category = workflow.group;
    if (routes[category] === undefined) {
      routes[category] = { workflows: [] };
    }

    routes[category].workflows = [
      ...routes[category].workflows,
      ...workflow.routes.map(route => {
        return {
          displayName: getDisplayName(workflow, route, " -"),
          path: `${workflow.path}/${route.path}`,
          trending: route.trending,
        };
      }),
    ];
  });
  return routes;
};

const sortedGroupings = (workflows: Workflow[]): string[] => {
  return Object.keys(routesByGrouping(workflows)).sort();
};

export interface SearchIndex {
  category: string;
  label: string;
  path: string;
}

const searchIndexes = (workflows: Workflow[]): SearchIndex[] => {
  let indexOptions = [];
  workflows.forEach(workflow => {
    const category = workflow.group;
    indexOptions = [
      ...indexOptions,
      ...workflow.routes.map(route => {
        const label = route.displayName
          ? `${workflow.displayName} - ${route.displayName}`
          : workflow.displayName;
        return {
          category,
          label,
          path: `${workflow.path}/${route.path}`,
        };
      }),
    ];
  });
  return indexOptions;
};

export { routesByGrouping, searchIndexes, sortedGroupings, workflowsByTrending };
